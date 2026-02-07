package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/application/query"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	routePath            = "/api/v1/users"
	ValidationErrMessage = "refresh token validation failed"
	InValidErrMessage    = "Invalid request body"
	ModulePrefix         = "UserModule - Restful"
)

type RefreshTokenHandler struct {
	Handler ports.QueryHandler[query.RefreshTokenQueryV1, *query.RefreshTokenQueryResponseV1]
}

func NewRefreshTokenHandler(
	handler ports.QueryHandler[query.RefreshTokenQueryV1, *query.RefreshTokenQueryResponseV1],
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		Handler: handler,
	}
}

// RefreshToken
//
//	@Summary	Refresh User Token
//	@Schemes
//	@Description	Refreshes the authentication token for a user using a valid refresh token.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body	RefreshTokenRequestDto	true	"Refresh Token Request DTO"
//	@Router			/refresh_token [post]
func (h RefreshTokenHandler) RefreshToken(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name) // go-arch
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	// todo context should be passed from gin context
	ctx, span := tracer.Start(c.Request.Context(), path)
	defer span.End()

	// Decode the request body into RefreshTokenRequestDto
	refreshTokenDto, decodeErr := helper.Decode[RefreshTokenRequestDto](c)
	if decodeErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", InValidErrMessage))
			span.SetStatus(codes.Error, InValidErrMessage)
			span.RecordError(decodeErr)
		}

		helper.ErrorResponse(c, InValidErrMessage, decodeErr, http.StatusBadRequest)
		return
	}

	// Validate the request body
	if validationErr := helper.Get().Struct(refreshTokenDto); validationErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ValidationErrMessage))
			span.SetStatus(codes.Error, ValidationErrMessage)
			span.RecordError(validationErr)
		}

		// If validation fails, extract the validation errors and return a validation error response
		helper.ValidationErrorResponse(c, ValidationErrMessage, validationErr)
		return
	}

	result, handlerErr := h.Handler.Handle(ctx, query.RefreshTokenQueryV1{
		RefreshToken: refreshTokenDto.RefreshToken,
		Platform:     helper.Platform(c),
	})
	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", handlerErr.Error()))
			span.SetStatus(codes.Error, handlerErr.Error())
			span.RecordError(handlerErr)
		}

		helper.UnauthorizedErrorResponse(
			c,
			handlerErr.Error(),
			handlerErr,
			http.StatusUnauthorized,
		)
		return
	}

	responsePayload := helper.Response[RefreshTokenResponseDto]{
		Data: &RefreshTokenResponseDto{
			UserID:       result.UserID,
			AccessToken:  result.AccessToken,
			RefreshToken: result.RefreshToken,
		},
		Links: []helper.Link{},
	}

	span.SetAttributes(attribute.String("users.id", result.UserID))

	helper.SuccessResponse(
		c,
		"refresh token updated successfully",
		responsePayload,
		http.StatusOK,
	)
}
