package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	routePath            = "/api/v1/users"
	ValidationErrMessage = "user validation request body does not validate"
	InValidErrMessage    = "Invalid request body"
	ModulePrefix         = "UserModule - Restful"
)

type RegisterUserHandler struct {
	Handler ports.CommandHandler[commands.RegisterUserCommandV1]
}

func NewRegisterUserHandler(
	handler ports.CommandHandler[commands.RegisterUserCommandV1],
) *RegisterUserHandler {
	return &RegisterUserHandler{
		Handler: handler,
	}
}

// Store It creates a new user
//
//	@Summary	Signup User
//	@Schemes
//	@Description	It is a method to create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body	RegisterUserRequestDto	true	"the user registration request body"
//	@Router			/auth/signup [post]
func (h RegisterUserHandler) Store(c *gin.Context) {
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

	// Decode the request body into RegisterUserRequestDto
	registerUserDto, decodeErr := helper.Decode[RegisterUserRequestDto](c)
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
	if validationErr := helper.Get().Struct(registerUserDto); validationErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ValidationErrMessage))
			span.SetStatus(codes.Error, ValidationErrMessage)
			span.RecordError(validationErr)
		}

		// If validation fails, extract the validation errors and return a validation error response
		helper.ValidationErrorResponse(c, ValidationErrMessage, validationErr)
		return
	}

	newID := uuid.NewID()

	handlerErr := h.Handler.Handle(ctx, commands.RegisterUserCommandV1{
		ID:       newID,
		Name:     registerUserDto.Name,
		Email:    registerUserDto.Email,
		Password: registerUserDto.Password,
	})
	if handlerErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", "user registration failed"))
			span.SetStatus(codes.Error, "user registration failed")
			span.RecordError(handlerErr)
		}

		helper.ExistFoundErrorResponse(
			c,
			"user registration failed",
			handlerErr,
			http.StatusInternalServerError,
		)
		return
	}

	responsePayload := helper.Response[RegisterUserResponseDto]{
		Links: []helper.Link{
			helper.AddHateoas(
				"self", // todo it can be "me" endpoint in future
				fmt.Sprintf("%s/%s", routePath, newID),
				http.MethodGet,
				""),
			helper.AddHateoas(
				"store",
				fmt.Sprintf("%s/", routePath),
				http.MethodPost,
				"/api/v1/schemas/users/register",
			),
			helper.AddHateoas(
				"update",
				fmt.Sprintf("%s/%s", routePath, newID),
				http.MethodPut,
				"/api/v1/schemas/users/update",
			),
			helper.AddHateoas(
				"delete",
				fmt.Sprintf("%s/%s", routePath, newID),
				http.MethodDelete,
				"",
			),
		},
	}

	span.SetAttributes(attribute.String("users.id", newID))

	helper.SuccessResponse(
		c,
		"User is registered successfully",
		responsePayload,
		http.StatusCreated,
	)
}
