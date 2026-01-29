package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	query "github.com/racibaz/go-arch/internal/modules/user/features/login/v1/application/queries"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	routePath = "/api/v1/auth/login"

	InValidErrMessage  = "Invalid request body"
	NotFoundErrMessage = "the user not found"
	ModulePrefix       = "UserModule - Restful"
)

type LoginHandler struct {
	Handler ports2.QueryHandler[query.LoginQueryV1, *query.LoginQueryResponse]
}

func NewLoginHandler(
	handler ports2.QueryHandler[query.LoginQueryV1, *query.LoginQueryResponse],
) *LoginHandler {
	return &LoginHandler{
		Handler: handler,
	}
}

func (h LoginHandler) Login(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(c, path)
	defer span.End()

	// Decode request body
	loginRequestDto, decodeErr := helper.Decode[LoginRequestDto](c)
	// Handle decode error
	if decodeErr != nil {
		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", InValidErrMessage))
			span.SetStatus(codes.Error, InValidErrMessage)
			span.RecordError(decodeErr)
		}

		helper.ValidationErrorResponse(c, InValidErrMessage, decodeErr)
		return
	}

	// Call the handler
	loginResult, handlerErr := h.Handler.Handle(ctx, query.LoginQueryV1{
		Email:    loginRequestDto.Email,
		Password: loginRequestDto.Password,
		Platform: helper.Platform(c), // Get platform from request
	})

	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", NotFoundErrMessage))
			span.SetStatus(codes.Error, NotFoundErrMessage)
			span.RecordError(handlerErr)
		}

		helper.ErrorResponse(c, NotFoundErrMessage, handlerErr, http.StatusInternalServerError)
		return
	}

	responsePayload := helper.Response[LoginResponseDto]{
		Data: &LoginResponseDto{
			AccessToken:  loginResult.AccessToken,
			RefreshToken: loginResult.RefreshToken,
			UserID:       loginResult.UserID,
		},
		Links: []helper.Link{
			helper.AddHateoas(
				"self",
				fmt.Sprintf("%s", routePath),
				http.MethodPost,
				"",
			),
		},
	}

	helper.SuccessResponse(c, "User LoggedIn", responsePayload, http.StatusOK)
}
