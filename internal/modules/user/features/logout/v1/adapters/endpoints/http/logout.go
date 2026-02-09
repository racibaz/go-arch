package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	commands "github.com/racibaz/go-arch/internal/modules/user/features/logout/v1/application/commands"
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
	routePath = "/api/v1/auth"

	NotFoundErrMessage = "the user not found"
	ModulePrefix       = "UserModule - Restful"
)

type LogoutHandler struct {
	Handler ports2.CommandHandler[commands.LogoutCommandV1]
}

func NewLogoutHandler(
	handler ports2.CommandHandler[commands.LogoutCommandV1],
) *LogoutHandler {
	return &LogoutHandler{
		Handler: handler,
	}
}

func (h LogoutHandler) Logout(c *gin.Context) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"%s - %s - %s",
		ModulePrefix,
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(c, path)
	defer span.End()

	// Call the handler
	handlerErr := h.Handler.Handle(ctx, commands.LogoutCommandV1{
		UserID:   c.GetString(helper.CtxUserID),
		Platform: c.GetString(helper.CtxPlatform),
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

	responsePayload := helper.Response[LogoutResponseDto]{
		Data: nil,
		Links: []helper.Link{
			helper.AddHateoas(
				"self",
				fmt.Sprintf("%s/%s", routePath, "logout"),
				http.MethodPost,
				"",
			),
			helper.AddHateoas(
				"login",
				fmt.Sprintf("%s/%s", routePath, "login"),
				http.MethodPost,
				"/api/v1/schemas/auth/login",
			),
		},
	}

	helper.SuccessResponse(c, "User Logged out", responsePayload, http.StatusOK)
}
