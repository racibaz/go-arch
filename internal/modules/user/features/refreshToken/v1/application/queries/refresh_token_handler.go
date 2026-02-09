package queries

import (
	"context"
	"fmt"
	"time"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type RefreshHandler struct {
	userRepository   ports.UserRepository
	logger           logger.Logger
	messagePublisher messaging.UserMessagePublisher
	tracer           trace.Tracer
}

var _ applicationPorts.QueryHandler[RefreshTokenQueryV1, *RefreshTokenQueryResponseV1] = (*RefreshHandler)(
	nil,
)

func NewRefreshHandler(
	userRepository ports.UserRepository,
	logger logger.Logger,
	messagePublisher messaging.UserMessagePublisher,
) *RefreshHandler {
	return &RefreshHandler{
		userRepository:   userRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("RefreshHandler"),
	}
}

func (h RefreshHandler) Handle(
	ctx context.Context,
	cmd RefreshTokenQueryV1,
) (*RefreshTokenQueryResponseV1, error) {
	ctx, span := h.tracer.Start(ctx, "RefreshHandler - Handler")
	defer span.End()

	existingUser, getUserByRefreshTokenErr := h.getUserByRefreshToken(
		ctx,
		cmd.RefreshToken,
		cmd.Platform,
	)

	if getUserByRefreshTokenErr != nil {
		h.logger.Error("user not found by refresh token: %v", getUserByRefreshTokenErr.Error())
		return nil, getUserByRefreshTokenErr
	}

	_, expirationTimeErr := h.checkRefreshTokenExpireTime(existingUser, cmd.Platform)
	if expirationTimeErr != nil {
		return nil, fmt.Errorf(
			"failed to check refresh token expiration time: %v",
			expirationTimeErr,
		)
	}

	accessToken, generateJwtErr := helper.GenerateJWT(
		existingUser.ID(),
		existingUser.Name,
		cmd.Platform,
	)
	if generateJwtErr != nil {
		return nil, generateJwtErr
	}

	refreshToken, refreshTokenErr := helper.GenerateRefreshToken()
	if refreshTokenErr != nil {
		return nil, fmt.Errorf(
			"login failed - could not generate refresh token: %w",
			refreshTokenErr,
		)
	}

	updateRefreshTokenErr := h.updateRefreshToken(
		ctx,
		existingUser.ID(),
		refreshToken,
		cmd.Platform,
	)
	if updateRefreshTokenErr != nil {
		return nil, fmt.Errorf(
			"login failed - could not update user refresh token: %w",
			updateRefreshTokenErr,
		)
	}

	// Refresh the existing user's refresh token in the current context
	if messageErr := h.messagePublisher.PublishUserRefreshTokenUpdated(ctx, existingUser); messageErr != nil {
		return nil, fmt.Errorf("failed refresh token updated event: %v", messageErr)
	}

	// Log successful login
	h.logger.Info("Refresh Token Updated. Id: %s", existingUser.ID())

	return &RefreshTokenQueryResponseV1{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       existingUser.ID(),
	}, nil
}

func (h RefreshHandler) isRefreshTokenWebValid(refreshTokenAt *time.Time) bool {
	if refreshTokenAt == nil {
		return false
	}

	return refreshTokenAt.After(time.Now())
}

func (h RefreshHandler) checkRefreshTokenExpireTime(
	user *domain.User,
	platform string,
) (bool, error) {
	switch platform {
	case helper.PlatformWeb:
		return h.isRefreshTokenWebValid(user.RefreshTokenWebAt), nil

	case helper.PlatformMobile:
		return h.isRefreshTokenWebValid(user.RefreshTokenMobileAt), nil

	default:
		return false, fmt.Errorf("invalid platform: %s", platform)
	}
}

func (h RefreshHandler) getUserByRefreshToken(
	ctx context.Context,
	refreshToken, platform string,
) (*domain.User, error) {
	switch platform {
	case helper.PlatformWeb:
		return h.userRepository.GetUserByRefreshTokenAtWeb(ctx, refreshToken)
	case helper.PlatformMobile:
		return h.userRepository.GetUserByRefreshTokenAtMobile(ctx, refreshToken)
	default:
		return nil, fmt.Errorf("invalid platform: %s", platform)
	}
}

// todo it is code duplicated with login_handler.go, consider refactoring later
func (h RefreshHandler) updateRefreshToken(
	ctx context.Context,
	id, refreshToken, platform string,
) error {
	switch platform {
	case helper.PlatformWeb:
		return h.userRepository.UpdateWebUserRefreshToken(ctx, id, refreshToken)
	case helper.PlatformMobile:
		return h.userRepository.UpdateMobileUserRefreshToken(ctx, id, refreshToken)
	default:
		return fmt.Errorf("invalid platform: %s", platform)
	}
}
