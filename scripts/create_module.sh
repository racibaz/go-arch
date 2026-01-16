#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./scripts/create_module.sh <module_name>"
    exit 1
fi

MODULE_NAME=$1

PASCAL_MODULE_NAME=$(echo "$MODULE_NAME" | awk -F'_' '{for(i=1;i<=NF;i++){ $i=toupper(substr($i,1,1)) substr($i,2)} }1' OFS='')

CAMEL_MODULE_NAME="$(tr '[:upper:]' '[:lower:]' <<< ${PASCAL_MODULE_NAME:0:1})${PASCAL_MODULE_NAME:1}"

MODULE_DIR="internal/modules/$MODULE_NAME"

if [ -d "$MODULE_DIR" ]; then
    echo "❌ Module already exists: $MODULE_NAME"
    exit 1
fi


echo "🚀 Creating module: $MODULE_NAME -> $PASCAL_MODULE_NAME"
echo "📌 PascalCase: $PASCAL_MODULE_NAME | camelCase: $CAMEL_MODULE_NAME"

mkdir -p internal/modules/$CAMEL_MODULE_NAME/application
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/commands
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/dtos
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/transformers
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints/http
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints/grpc
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints/grpc/proto
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/handlers
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/ports

mkdir -p internal/modules/$CAMEL_MODULE_NAME/domain
mkdir -p internal/modules/$CAMEL_MODULE_NAME/domain/ports

mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/messaging
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/messaging/rabbitmq
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/notification
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/notification/sms
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/entities
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/mappers
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/repositories
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/in_memory

mkdir -p internal/modules/$CAMEL_MODULE_NAME/logging
mkdir -p internal/modules/$CAMEL_MODULE_NAME/routes
# TODO: Add testing directories and mocks when needed

cat > internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints/http/create_${CAMEL_MODULE_NAME}.go << EOF
package http

import (
	"fmt"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/features/creating${CAMEL_MODULE_NAME}/v1/dtos"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/features/creating${CAMEL_MODULE_NAME}/v1/commands"

	ports "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	routePath = "/api/v1/${CAMEL_MODULE_NAME}"

	ValidationErrMessage = "${CAMEL_MODULE_NAME} validation request body does not validate"
	InValidErrMessage    = "Invalid request body"
	NotFoundErrMessage   = "Not found a record"
	ParseErrMessage      = "The Id does not parse correctly"
	ModulePrefix         = "PostModule - Restful"
)


// Create${PASCAL_MODULE_NAME}ResponseDto
// @Description Create${PASCAL_MODULE_NAME}ResponseDto is a data transfer object for reporting the details of a created ${CAMEL_MODULE_NAME}
type Create${PASCAL_MODULE_NAME}ResponseDto struct {
	// @Description ID is the id of the ${CAMEL_MODULE_NAME}
	ID string \`json:"id"\`
}

// Get${PASCAL_MODULE_NAME}ResponseDto
// @Description Get${PASCAL_MODULE_NAME}ResponseDto is a data transfer object for reporting the details of a ${CAMEL_MODULE_NAME}
type Get${PASCAL_MODULE_NAME}ResponseDto struct {
	// @Description ID is the id of the ${CAMEL_MODULE_NAME}
	ID          string    \`json:"id"\`
	UserID      string    \`json:"userId"\`
	Title       string    \`json:"title"\`
	Description string    \`json:"description"\`
	Content     string    \`json:"content"\`
	Status      int       \`json:"status"\`
	CreatedAt   time.Time \`json:"createdAt"\`
	UpdatedAt   time.Time \`json:"updatedAt"\`
}

// Create${PASCAL_MODULE_NAME}RequestDto
// @Description Create${PASCAL_MODULE_NAME}RequestDto is a data transfer object for creating a ${CAMEL_MODULE_NAME}
type Create${PASCAL_MODULE_NAME}RequestDto struct {
	// @Description UserID is the ID of the user creating the ${CAMEL_MODULE_NAME}
	UserID string \`json:"userId" validate:"required,uuid"\`
	// @Description Title of the ${CAMEL_MODULE_NAME}
	Title string \`json:"title" validate:"required,min=10"\`
	// @Description Description of the ${CAMEL_MODULE_NAME}
	Description string \`json:"description" validate:"required,min=10"\`
	// @Description Content of the ${CAMEL_MODULE_NAME}
	Content string \`json:"content" validate:"required,min=10"\`
}


type ${PASCAL_MODULE_NAME}Handler struct {
	Create${PASCAL_MODULE_NAME}Handler *commands.Create${PASCAL_MODULE_NAME}Handler
	Repository                   ports.${PASCAL_MODULE_NAME}Repository
}

func New${PASCAL_MODULE_NAME}Handler(createHandler *commands.Create${PASCAL_MODULE_NAME}Handler, repository ports.${PASCAL_MODULE_NAME}Repository) *${PASCAL_MODULE_NAME}Handler {
	return &${PASCAL_MODULE_NAME}Handler{
		Create${PASCAL_MODULE_NAME}Handler: createHandler,
		Repository:                   repository,
	}
}



//	@BasePath	/api/v1

// Store ${PASCAL_MODULE_NAME}Store Store is a method to create a new ${CAMEL_MODULE_NAME}
//
//	@Summary	${CAMEL_MODULE_NAME} store
//	@Schemes
//	@Description	It is a method to create a new ${CAMEL_MODULE_NAME}
//	@Tags			${CAMEL_MODULE_NAME}s
//	@Accept			json
//	@Produce		json
//	@Param			post	body		Create${PASCAL_MODULE_NAME}RequestDto	true	"Create ${PASCAL_MODULE_NAME} Request DTO"
//	@Success		201		{object}	Create${PASCAL_MODULE_NAME}ResponseDto	"${PASCAL_MODULE_NAME} created successfully"
//	@Failure		400		{object}	errors.AppError					"Invalid request body"
//	@Router			/${CAMEL_MODULE_NAME}s [post]
func (handler *${PASCAL_MODULE_NAME}Handler) Store(c *gin.Context) {

	tracer := otel.Tracer(config.Get().App.Name)
  path := fmt.Sprintf("${PASCAL_MODULE_NAME}Module - Restful - %s - %s", helper.StructName(handler), helper.CurrentFuncName())
	ctx, span := tracer.Start(c.Request.Context(), path)
	defer span.End()

	// Decode the request body into Create${PASCAL_MODULE_NAME}RequestDto
	create${PASCAL_MODULE_NAME}RequestDto, decodeErr := helper.Decode[dtos.Create${PASCAL_MODULE_NAME}RequestDto](c)
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
	if validationErr := helper.Get().Struct(create${PASCAL_MODULE_NAME}RequestDto); validationErr != nil {

		if span := trace.SpanFromContext(ctx); span != nil {
			span.SetAttributes(attribute.String("error", ValidationErrMessage))
			span.SetStatus(codes.Error, ValidationErrMessage)
			span.RecordError(validationErr)
		}

		// If validation fails, extract the validation errors and return a validation error response
		helper.ValidationErrorResponse(c, ValidationErrMessage, validationErr)
	}


	newUuid := uuid.NewID()

	command := commands.Create${PASCAL_MODULE_NAME}CommandV1{
		ID:          newUuid,
		UserID:      create${PASCAL_MODULE_NAME}RequestDto.UserID,
		Title:       create${PASCAL_MODULE_NAME}RequestDto.Title,
		Description: create${PASCAL_MODULE_NAME}RequestDto.Description,
		Content:     create${PASCAL_MODULE_NAME}RequestDto.Content,
		Status:      domain.${PASCAL_MODULE_NAME}StatusDraft,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = handler.Create${PASCAL_MODULE_NAME}Handler.Handle(ctx, command)

	if err != nil {
		span.SetAttributes(attribute.String("error", "${PASCAL_MODULE_NAME} create failed"))
		span.SetStatus(codes.Error, "${PASCAL_MODULE_NAME} create failed")
		helper.ErrorResponse(c, "${CAMEL_MODULE_NAME} create failed", err, http.StatusInternalServerError)
		return
	}

	responseData := Create${PASCAL_MODULE_NAME}ResponseDto{
		ID:       create${PASCAL_MODULE_NAME}RequestDto.ID,
	}

	span.SetAttributes(attribute.String("${CAMEL_MODULE_NAME}.id", newUuid))

	helper.SuccessResponse(c, "${PASCAL_MODULE_NAME} created successfully", responseData, http.StatusCreated)
}


EOF


cat > internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/commands/create_${CAMEL_MODULE_NAME}_command.go << EOF
package commands

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
)

type Create${PASCAL_MODULE_NAME}CommandV1 struct {
	ID          string // Unique identifier for the ${CAMEL_MODULE_NAME}
	UserID      string
	Title       string
	Description string
	Content     string
	Status      domain.${PASCAL_MODULE_NAME}Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/commands/create_${CAMEL_MODULE_NAME}_handler.go << EOF
package commands

import (
	"context"
	"fmt"
	"time"

	applicationPorts "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Create${PASCAL_MODULE_NAME}Handler handles the creation of new ${CAMEL_MODULE_NAME}s.
type Create${PASCAL_MODULE_NAME}Handler struct {
	${PASCAL_MODULE_NAME}Repository   ports.${PASCAL_MODULE_NAME}Repository
	logger           logger.Logger
	messagePublisher messaging.MessagePublisher
	tracer           trace.Tracer
}

// Ensure Create${PASCAL_MODULE_NAME}Handler implements the CommandHandler interface
var _ applicationPorts.CommandHandler[Create${PASCAL_MODULE_NAME}CommandV1] = (*Create${PASCAL_MODULE_NAME}Handler)(nil)

// NewCreate${PASCAL_MODULE_NAME}Handler initializes a new Create${PASCAL_MODULE_NAME}Handler with the provided ${PASCAL_MODULE_NAME}Repository.
func NewCreate${PASCAL_MODULE_NAME}Handler(
	${CAMEL_MODULE_NAME}Repository ports.${PASCAL_MODULE_NAME}Repository,
	logger logger.Logger,
	messagePublisher messaging.MessagePublisher,
) *Create${PASCAL_MODULE_NAME}Handler {
	return &Create${PASCAL_MODULE_NAME}Handler{
		${PASCAL_MODULE_NAME}Repository:   ${CAMEL_MODULE_NAME}Repository,
		logger:           logger,
		messagePublisher: messagePublisher,
		tracer:           otel.Tracer("Create${PASCAL_MODULE_NAME}Handler"),
	}
}

// Handle processes the Create${PASCAL_MODULE_NAME}CommandV1 to create a new ${CAMEL_MODULE_NAME}.
func (h Create${PASCAL_MODULE_NAME}Handler) Handle(ctx context.Context, cmd Create${PASCAL_MODULE_NAME}CommandV1) error {
	ctx, span := h.tracer.Start(ctx, "Create${PASCAL_MODULE_NAME} - Handler")
	defer span.End()

	// Create a new ${CAMEL_MODULE_NAME} using the factory
	${CAMEL_MODULE_NAME}, _ := domain.Create(
		cmd.ID,
		cmd.UserID,
		cmd.Title,
		cmd.Description,
		cmd.Content,
		cmd.Status,
		time.Now(),
		time.Now(),
	)

	// Check if the ${CAMEL_MODULE_NAME} exists in db
	isExists, err := h.${PASCAL_MODULE_NAME}Repository.IsExists(ctx, ${CAMEL_MODULE_NAME}.Title, ${CAMEL_MODULE_NAME}.Description)
	if err != nil {
		h.logger.Error("Error saving ${CAMEL_MODULE_NAME}: %v", err.Error())
		return fmt.Errorf("error checking if ${CAMEL_MODULE_NAME} exists: %v", err)
	}

	// If the ${CAMEL_MODULE_NAME} already exists, return an error
	if isExists {
		h.logger.Info(
			"${PASCAL_MODULE_NAME} already exists with title: %s and description: %s",
			${CAMEL_MODULE_NAME}.Title,
			${CAMEL_MODULE_NAME}.Description,
		)
		return domain.ErrAlreadyExists
	}

	// Save the new ${CAMEL_MODULE_NAME} to the repository
	savingErr := h.${PASCAL_MODULE_NAME}Repository.Save(ctx, ${CAMEL_MODULE_NAME})

	if savingErr != nil {
		h.logger.Error("Error saving ${CAMEL_MODULE_NAME}: %v", savingErr)
		return savingErr
	}

	// Publish an event indicating that a new ${CAMEL_MODULE_NAME} has been created
	if h.messagePublisher != nil {
		// TODO: Implement message publishing when publisher is available
		// if messageErr := h.messagePublisher.Publish${PASCAL_MODULE_NAME}Created(ctx, ${CAMEL_MODULE_NAME}); messageErr != nil {
		//	return fmt.Errorf("failed to publish the ${CAMEL_MODULE_NAME} created event: %v", messageErr)
		// }
	}

	h.logger.Info("${PASCAL_MODULE_NAME} created successfully with ID: %s", ${CAMEL_MODULE_NAME}.ID())

	return nil
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/entities/${CAMEL_MODULE_NAME}_entity.go << EOF
package entities

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
)

type ${PASCAL_MODULE_NAME} struct {
	ID          string                    \`gorm:"primaryKey;type:uuid;default:gen_random_uuid()"\`
	UserID      string                    \`gorm:"not null"\`
	Title       string                    \`gorm:"not null;size:255"\`
	Description string                    \`gorm:"not null;size:500"\`
	Content     string                    \`gorm:"not null;type:text"\`
	Status      domain.${PASCAL_MODULE_NAME}Status \`gorm:"not null;type:varchar(20);default:'draft'"\`
	CreatedAt   time.Time                 \`gorm:"autoCreateTime"\`
	UpdatedAt   time.Time                 \`gorm:"autoUpdateTime"\`
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/mappers/${CAMEL_MODULE_NAME}_mapper.go << EOF
package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	entity "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/es"
)

func ToDomain(${CAMEL_MODULE_NAME}Entity *entity.${PASCAL_MODULE_NAME}) *domain.${PASCAL_MODULE_NAME} {
	return &domain.${PASCAL_MODULE_NAME}{
		Aggregate:   es.NewAggregate(${CAMEL_MODULE_NAME}Entity.ID, domain.${PASCAL_MODULE_NAME}Aggregate),
		UserID:      ${CAMEL_MODULE_NAME}Entity.UserID,
		Title:       ${CAMEL_MODULE_NAME}Entity.Title,
		Description: ${CAMEL_MODULE_NAME}Entity.Description,
		Content:     ${CAMEL_MODULE_NAME}Entity.Content,
		Status:      ${CAMEL_MODULE_NAME}Entity.Status,
		CreatedAt:   ${CAMEL_MODULE_NAME}Entity.CreatedAt,
		UpdatedAt:   ${CAMEL_MODULE_NAME}Entity.UpdatedAt,
	}
}

func ToPersistence(${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) *entity.${PASCAL_MODULE_NAME} {
	return &entity.${PASCAL_MODULE_NAME}{
		ID:          ${CAMEL_MODULE_NAME}.ID(),
		UserID:      ${CAMEL_MODULE_NAME}.UserID,
		Title:       ${CAMEL_MODULE_NAME}.Title,
		Description: ${CAMEL_MODULE_NAME}.Description,
		Content:     ${CAMEL_MODULE_NAME}.Content,
		Status:      ${CAMEL_MODULE_NAME}.Status,
		CreatedAt:   ${CAMEL_MODULE_NAME}.CreatedAt,
		UpdatedAt:   ${CAMEL_MODULE_NAME}.UpdatedAt,
	}
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/repositories/${CAMEL_MODULE_NAME}_repository.go << EOF
package repositories

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/infrastructure/persistence/gorm/entities"
	${CAMEL_MODULE_NAME}Mapper "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/infrastructure/persistence/gorm/mappers"
	"github.com/racibaz/go-arch/pkg/database"
	"gorm.io/gorm"
	"sync"
)

// Gorm${PASCAL_MODULE_NAME}Repository Secondary adapter: PostgreSQL implementation
type Gorm${PASCAL_MODULE_NAME}Repository struct {
	DB *gorm.DB
	sync.Mutex
}

var _ ports.${PASCAL_MODULE_NAME}Repository = (*Gorm${PASCAL_MODULE_NAME}Repository)(nil)

func NewGorm${PASCAL_MODULE_NAME}Repository() *Gorm${PASCAL_MODULE_NAME}Repository {
	return &Gorm${PASCAL_MODULE_NAME}Repository{
		DB: database.Connection(),
	}
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) Save(ctx context.Context, ${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) error {
	persistenceModel := ${CAMEL_MODULE_NAME}Mapper.ToPersistence(${CAMEL_MODULE_NAME})

	err := repo.DB.WithContext(ctx).Create(persistenceModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) GetByID(ctx context.Context, id string) (*domain.${PASCAL_MODULE_NAME}, error) {
	var entity entities.${PASCAL_MODULE_NAME}

	if err := repo.DB.WithContext(ctx).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}

	return ${CAMEL_MODULE_NAME}Mapper.ToDomain(&entity), nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) Update(ctx context.Context, ${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) error {
	persistenceModel := ${CAMEL_MODULE_NAME}Mapper.ToPersistence(${CAMEL_MODULE_NAME})

	err := repo.DB.WithContext(ctx).Save(persistenceModel).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) Delete(ctx context.Context, id string) error {
	err := repo.DB.WithContext(ctx).Where("id = ?", id).Delete(&entities.${PASCAL_MODULE_NAME}{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) List(ctx context.Context) ([]*domain.${PASCAL_MODULE_NAME}, error) {
	var entities []entities.${PASCAL_MODULE_NAME}

	if err := repo.DB.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}

	var ${CAMEL_MODULE_NAME}s []*domain.${PASCAL_MODULE_NAME}
	for _, entity := range entities {
		${CAMEL_MODULE_NAME}s = append(${CAMEL_MODULE_NAME}s, ${CAMEL_MODULE_NAME}Mapper.ToDomain(&entity))
	}

	return ${CAMEL_MODULE_NAME}s, nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) IsExists(ctx context.Context, title, description string) (bool, error) {
	var count int64

	err := repo.DB.WithContext(ctx).Model(&entities.${PASCAL_MODULE_NAME}{}).
		Where("title = ? AND description = ?", title, description).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/domain/${CAMEL_MODULE_NAME}_status.go << EOF
package domain

import "errors"

type ${PASCAL_MODULE_NAME}Status string

const (
	${PASCAL_MODULE_NAME}StatusDraft    ${PASCAL_MODULE_NAME}Status = "draft"
	${PASCAL_MODULE_NAME}StatusPublished ${PASCAL_MODULE_NAME}Status = "published"
	${PASCAL_MODULE_NAME}StatusArchived  ${PASCAL_MODULE_NAME}Status = "archived"
)

var ErrInvalidStatus = errors.New("status is not valid")

func (${PASCAL_MODULE_NAME}Status) IsValid() error {
	switch ${PASCAL_MODULE_NAME}Status {
	case ${PASCAL_MODULE_NAME}StatusDraft, ${PASCAL_MODULE_NAME}StatusPublished, ${PASCAL_MODULE_NAME}StatusArchived:
		return nil
	default:
		return ErrInvalidStatus
	}
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/domain/${CAMEL_MODULE_NAME}.go << EOF
package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/racibaz/go-arch/pkg/es"
)

const (
	${PASCAL_MODULE_NAME}Aggregate = "${CAMEL_MODULE_NAME}s.${PASCAL_MODULE_NAME}"
)

var (
	TitleMinLength       = 10
	DescriptionMinLength = 10
	ContentMinLength     = 10
)

var (
	ErrNotFound          = errors.New("the ${CAMEL_MODULE_NAME} was not found")
	ErrAlreadyExists     = errors.New("the ${CAMEL_MODULE_NAME} already exists")
	ErrEmptyId           = errors.New("id cannot be empty")
	ErrEmptyUserId       = errors.New("user id cannot be empty")
	ErrMinTitleLength = errors.New(
		fmt.Sprintf("title must be at least %d characters long", TitleMinLength),
	)
	ErrMinDescriptionLength = errors.New(
		fmt.Sprintf("description must be at least %d characters long", DescriptionMinLength),
	)
	ErrMinContentLength = errors.New(
		fmt.Sprintf("content must be at least %d characters long", ContentMinLength),
	)
)

type ${PASCAL_MODULE_NAME} struct {
	es.Aggregate
	UserID      string
	Title       string
	Description string
	Content     string
	Status      ${PASCAL_MODULE_NAME}Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (p *${PASCAL_MODULE_NAME}) validate() error {
	if p.ID() == "" {
		return ErrEmptyId
	}
	if p.UserID == "" {
		return ErrEmptyUserId
	}
	if len(strings.TrimSpace(p.Title)) < TitleMinLength {
		return ErrMinTitleLength
	}
	if len(strings.TrimSpace(p.Description)) < DescriptionMinLength {
		return ErrMinDescriptionLength
	}
	if len(strings.TrimSpace(p.Content)) < ContentMinLength {
		return ErrMinContentLength
	}
	if err := p.Status.IsValid(); err != nil {
		return err
	}
	return nil
}

// Create This factory method creates a new ${PASCAL_MODULE_NAME} with default values if you want.
func Create(
	id, userID, title, description, content string,
	status ${PASCAL_MODULE_NAME}Status,
	createdAt, updatedAt time.Time,
) (*${PASCAL_MODULE_NAME}, error) {
	${CAMEL_MODULE_NAME} := &${PASCAL_MODULE_NAME}{
		Aggregate:   es.NewAggregate(id, ${PASCAL_MODULE_NAME}Aggregate),
		UserID:      userID,
		Title:       title,
		Description: description,
		Content:     content,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	// validate the ${CAMEL_MODULE_NAME} before returning it
	err := ${CAMEL_MODULE_NAME}.validate()
	if err != nil {
		return nil, err
	}

	return ${CAMEL_MODULE_NAME}, nil
}

func (p *${PASCAL_MODULE_NAME}) Delete() {
	// todo implement me
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/domain/ports/${CAMEL_MODULE_NAME}_repository.go << EOF
package ports

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
)

// ${PASCAL_MODULE_NAME}Repository Secondary port: ${PASCAL_MODULE_NAME}Repository interface
type ${PASCAL_MODULE_NAME}Repository interface {
	Save(ctx context.Context, ${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) error
	GetByID(ctx context.Context, id string) (*domain.${PASCAL_MODULE_NAME}, error)
	Update(ctx context.Context, ${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*domain.${PASCAL_MODULE_NAME}, error)
	IsExists(ctx context.Context, title, description string) (bool, error)
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/routes/routes.go << EOF
package routes

import (
	"sync"

	"github.com/gin-gonic/gin"
	module "github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/commands"
	creatingV1Endpoint "github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/endpoints/grpc"
	getByIdV1Endpoint "github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/features/getting${CAMEL_MODULE_NAME}byid/v1/endpoints"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/features/getting${CAMEL_MODULE_NAME}byid/v1/query"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/application/handlers"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/infrastructure/messaging/rabbitmq"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/infrastructure/notification/sms"
	gormRepo "github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/internal/modules/$CAMEL_MODULE_NAME/logging"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
	rabbitmqConn "github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	googleGrpc "google.golang.org/grpc"
)

var (
	${CAMEL_MODULE_NAME}ModuleInstance *module.${PASCAL_MODULE_NAME}Module
	once                              sync.Once
)

func Build${PASCAL_MODULE_NAME}Module() *module.${PASCAL_MODULE_NAME}Module {
	// Return existing instance if already created
	if ${CAMEL_MODULE_NAME}ModuleInstance != nil {
		return ${CAMEL_MODULE_NAME}ModuleInstance
	}

	// Create the instance only once
	once.Do(func() {
		// Use In-memory for persistence
		// repo := in_memory.NewGorm${PASCAL_MODULE_NAME}Repository()
		// Use GORM repository for persistence
		repository := gormRepo.NewGorm${PASCAL_MODULE_NAME}Repository()

		// Assuming NewZapLogger is a function that initializes a logger
		logger, _ := logger.NewZapLogger()

		domainDispatcher := ddd.NewEventDispatcher[ddd.AggregateEvent]()

		rabbitmqConn := rabbitmqConn.Connection()

		messagePublisher := rabbitmq.New${PASCAL_MODULE_NAME}MessagePublisher(rabbitmqConn, logger)
		/* todo we need to use processor in handler to publish events after transaction is committed
		for now we will use directly the publisher in the handler
		*/
		createCommandHandler := commands.NewCreate${PASCAL_MODULE_NAME}Handler(
			repository,
			logger,
			messagePublisher,
		)

		getQueryHandler := query.NewGet${PASCAL_MODULE_NAME}Handler(repository, logger)

		notificationAdapter := sms.NewTwilioSmsNotificationAdapter()

		notificationHandlers := logging.LogEventHandlerAccess[ddd.AggregateEvent](
			handlers.NewNotificationHandlers(notificationAdapter),
			"Notification", logger,
		)

		handlers.RegisterNotificationHandlers(notificationHandlers, domainDispatcher)

		${CAMEL_MODULE_NAME}ModuleInstance = module.New${PASCAL_MODULE_NAME}Module(
			repository,
			createCommandHandler,
			getQueryHandler,
			logger,
			notificationAdapter,
		)
	})

	return ${CAMEL_MODULE_NAME}ModuleInstance
}

func Routes(router *gin.Engine) {
	module := Build${PASCAL_MODULE_NAME}Module()

	// Collect here restful routes of your module.
	creatingV1Endpoint.MapHttpRoute(router, module.CommandHandler())
	getByIdV1Endpoint.MapHttpRoute(router, module.QueryHandler())
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {
	module := Build${PASCAL_MODULE_NAME}Module()

	// Collect here grpc routes of your module
	grpc.NewCreate${PASCAL_MODULE_NAME}Handler(grpcServer, module.CommandHandler())
}


EOF


cat > internal/modules/$CAMEL_MODULE_NAME/application/ports/command_handler.go << EOF
package ports

import "context"

type CommandHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) error
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/application/ports/query_handler.go << EOF
package ports

import "context"

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/application/features/creating${CAMEL_MODULE_NAME}/v1/dtos/dtos.go << EOF
package dtos

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
)

type Create${PASCAL_MODULE_NAME}Input struct {
	ID          string // Unique identifier for the ${CAMEL_MODULE_NAME}
	UserID      string
	Title       string
	Description string
	Content     string
	Status      domain.${PASCAL_MODULE_NAME}Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

EOF



cat > internal/modules/$CAMEL_MODULE_NAME/module_test.go << EOF
package module

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test${PASCAL_MODULE_NAME}Module(t *testing.T) {
	module := New${PASCAL_MODULE_NAME}Module()

	assert.NotNil(t, module)
	assert.NotNil(t, module.Repository())
	assert.NotNil(t, module.Create${PASCAL_MODULE_NAME}Handler())
	assert.NotNil(t, module.Logger())
}
EOF


cat > internal/modules/$CAMEL_MODULE_NAME/module.go << EOF
package module

import (
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/features/creating${CAMEL_MODULE_NAME}/v1/commands"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/features/getting${CAMEL_MODULE_NAME}byid/v1/query"
	ports "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	${CAMEL_MODULE_NAME}DomainPorts "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
)

// ${PASCAL_MODULE_NAME}Module encapsulates the components related to the ${PASCAL_MODULE_NAME} module.
type ${PASCAL_MODULE_NAME}Module struct {
	repository               ${CAMEL_MODULE_NAME}DomainPorts.${PASCAL_MODULE_NAME}Repository
	create${PASCAL_MODULE_NAME}CommandHandler ports.CommandHandler[commands.Create${PASCAL_MODULE_NAME}CommandV1]
	get${PASCAL_MODULE_NAME}QueryHandler      ports.QueryHandler[query.Get${PASCAL_MODULE_NAME}ByIdQuery, query.Get${PASCAL_MODULE_NAME}ByIdQueryResponse]
	logger                   logger.Logger
	notifier                 ${CAMEL_MODULE_NAME}DomainPorts.NotificationAdapter
}

// New${PASCAL_MODULE_NAME}Module initializes a new ${PASCAL_MODULE_NAME}Module with the provided components.
func New${PASCAL_MODULE_NAME}Module(
	repository ${CAMEL_MODULE_NAME}DomainPorts.${PASCAL_MODULE_NAME}Repository,
	create${PASCAL_MODULE_NAME}CommandHandler ports.CommandHandler[commands.Create${PASCAL_MODULE_NAME}CommandV1],
	get${PASCAL_MODULE_NAME}QueryHandler ports.QueryHandler[query.Get${PASCAL_MODULE_NAME}ByIdQuery, query.Get${PASCAL_MODULE_NAME}ByIdQueryResponse],
	logger logger.Logger,
	notifier ${CAMEL_MODULE_NAME}DomainPorts.NotificationAdapter,
) *${PASCAL_MODULE_NAME}Module {
	return &${PASCAL_MODULE_NAME}Module{
		repository:               repository,
		create${PASCAL_MODULE_NAME}CommandHandler: create${PASCAL_MODULE_NAME}CommandHandler,
		get${PASCAL_MODULE_NAME}QueryHandler:      get${PASCAL_MODULE_NAME}QueryHandler,
		logger:                   logger,
		notifier:                 notifier,
	}
}

func (m ${PASCAL_MODULE_NAME}Module) Repository() ${CAMEL_MODULE_NAME}DomainPorts.${PASCAL_MODULE_NAME}Repository {
	return m.repository
}

func (m ${PASCAL_MODULE_NAME}Module) CommandHandler() ports.CommandHandler[commands.Create${PASCAL_MODULE_NAME}CommandV1] {
	return m.create${PASCAL_MODULE_NAME}CommandHandler
}

func (m ${PASCAL_MODULE_NAME}Module) QueryHandler() ports.QueryHandler[query.Get${PASCAL_MODULE_NAME}ByIdQuery, query.Get${PASCAL_MODULE_NAME}ByIdQueryResponse] {
	return m.get${PASCAL_MODULE_NAME}QueryHandler
}

func (m ${PASCAL_MODULE_NAME}Module) Notifier() ${CAMEL_MODULE_NAME}DomainPorts.NotificationAdapter {
	return m.notifier
}

func (m ${PASCAL_MODULE_NAME}Module) Logger() logger.Logger {
	return m.logger
}

EOF

echo "✅ Module $PASCAL_MODULE_NAME created successfully!"
