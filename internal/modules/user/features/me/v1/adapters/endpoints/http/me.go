package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	"github.com/racibaz/go-arch/internal/modules/user/features/me/v1/application/queries"
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

type MeHttpHandler struct {
	Handler ports.QueryHandler[queries.MeQueryHandlerQuery, *queries.MeQueryHandlerResponse]
}

func NewMeHttpHandler(
	handler ports.QueryHandler[queries.MeQueryHandlerQuery, *queries.MeQueryHandlerResponse],
) *MeHttpHandler {
	return &MeHttpHandler{
		Handler: handler,
	}
}

// Me
//
//	@Summary	Get current user information
//	@Schemes
//	@Description			Get current user information
//	@Tags					users
//	@Accept					json
//	@Produce				json
//	@Param					user	body	MeRequestDto	true	"Me Request Object"
//	@Router					/users/me [get]
//	@Success				200	{object}	helper.Response[MeResponseDto]
//	@Security				BearerAuth
//	@SecurityDefinitions	BearerAuth
func (h MeHttpHandler) Me(c *gin.Context) {
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

	// Decode the request body into MeRequestDto
	meRequestDto, decodeErr := helper.Decode[MeRequestDto](c)
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
	if validationErr := helper.Get().Struct(meRequestDto); validationErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ValidationErrMessage))
			span.SetStatus(codes.Error, ValidationErrMessage)
			span.RecordError(validationErr)
		}

		// If validation fails, extract the validation errors and return a validation error response
		helper.ValidationErrorResponse(c, ValidationErrMessage, validationErr)
		return
	}

	result, handlerErr := h.Handler.Handle(ctx, queries.MeQueryHandlerQuery{
		RefreshToken: meRequestDto.RefreshToken,
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

	responsePayload := helper.Response[MeResponseDto]{
		Data: &MeResponseDto{
			Name:      result.Name,
			Email:     result.Email,
			Status:    domain.StatusToString[result.Status],
			CreatedAt: result.CreatedAt.String(),
		},
		Links: []helper.Link{
			helper.AddHateoas(
				"self",
				fmt.Sprintf("%s/%s", routePath, "/me"),
				http.MethodGet,
				""),
		},
	}

	span.SetAttributes(attribute.String("users.id", result.ID))

	helper.SuccessResponse(
		c,
		"user information retrieved successfully",
		responsePayload,
		http.StatusOK,
	)
}
