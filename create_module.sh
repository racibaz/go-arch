#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: ./create_module.sh <module_name>"
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
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/commands
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/queries
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/ports
mkdir -p internal/modules/$CAMEL_MODULE_NAME/application/dtos

mkdir -p internal/modules/$CAMEL_MODULE_NAME/domain
mkdir -p internal/modules/$CAMEL_MODULE_NAME/domain/ports

mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/entities
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/mappers
mkdir -p internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/repositories

mkdir -p internal/modules/$CAMEL_MODULE_NAME/presentation
mkdir -p internal/modules/$CAMEL_MODULE_NAME/presentation/http
mkdir -p internal/modules/$CAMEL_MODULE_NAME/presentation/routes

mkdir -p internal/modules/$CAMEL_MODULE_NAME/tests
mkdir -p internal/modules/$CAMEL_MODULE_NAME/query

cat > internal/modules/$CAMEL_MODULE_NAME/presentation/http/${CAMEL_MODULE_NAME}_controller.go << EOF
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/dtos"
	ports "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	validator "github.com/racibaz/go-arch/pkg/validator"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"net/http"
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
	ID string \`json:"id"\`
}

// Create${PASCAL_MODULE_NAME}RequestDto
// @Description Create${PASCAL_MODULE_NAME}RequestDto is a data transfer object for creating a ${CAMEL_MODULE_NAME}
type Create${PASCAL_MODULE_NAME}RequestDto struct {
	// @Description ID is the ID of the user creating the ${CAMEL_MODULE_NAME}
	ID string \`json:"id" validate:"required,uuid"\`
}


type ${PASCAL_MODULE_NAME}Controller struct {
	Service ports.${PASCAL_MODULE_NAME}Service
}

func New${PASCAL_MODULE_NAME}Controller(service ports.${PASCAL_MODULE_NAME}Service) *${PASCAL_MODULE_NAME}Controller {
	return &${PASCAL_MODULE_NAME}Controller{
		Service: service,
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
func (${CAMEL_MODULE_NAME}Controller ${PASCAL_MODULE_NAME}Controller) Store(c *gin.Context) {

	tracer := otel.Tracer("go-arch")
	ctx, span := tracer.Start(c.Request.Context(), "${PASCAL_MODULE_NAME}Module - Restful - ${PASCAL_MODULE_NAME}Controller - Store")
	defer span.End()

	create${PASCAL_MODULE_NAME}RequestDto, err := helper.Decode[Create${PASCAL_MODULE_NAME}RequestDto](c)

	if err != nil {
		helper.ErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}

	// Validate the request body
	if err := validator.Get().Struct(&create${PASCAL_MODULE_NAME}RequestDto); err != nil {
		// If validation fails, extract the validation errors
		c.JSON(
			http.StatusBadRequest,
			validator.NewValidationError(
				"${CAMEL_MODULE_NAME} validation request body does not validate",
				validator.ShowRegularValidationErrors(err).Errors,
			),
		)
		return
	}

	newUuid := uuid.NewID()

	err = ${CAMEL_MODULE_NAME}Controller.Service.Create${PASCAL_MODULE_NAME}(ctx, dto.Create${PASCAL_MODULE_NAME}Input{
		ID:          newUuid,
	})

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

// Show ${PASCAL_MODULE_NAME}GetById Show is a method to retrieve a ${CAMEL_MODULE_NAME} by its ID
//
//	@Summary	Get ${CAMEL_MODULE_NAME} by id
//	@Schemes
//	@Description	It is a method to retrieve a ${CAMEL_MODULE_NAME} by its ID
//	@Tags			${CAMEL_MODULE_NAME}s
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string		true	"${PASCAL_MODULE_NAME} ID"
//	@Success		200	{object}	Get${PASCAL_MODULE_NAME}ResponseDto	"${PASCAL_MODULE_NAME} retrieved successfully"
//	@Failure		404	{object}	errors.AppError		"Page not found"
//	@Router			/${CAMEL_MODULE_NAME}s/{id} [get]
func (${CAMEL_MODULE_NAME}Controller ${PASCAL_MODULE_NAME}Controller) Show(c *gin.Context) {

	tracer := otel.Tracer("go-arch")
	ctx, span := tracer.Start(c, "${PASCAL_MODULE_NAME}Module - Restful - ${PASCAL_MODULE_NAME}Controller - Show")
	defer span.End()

	${CAMEL_MODULE_NAME}ID := c.Param("id")

	result, err := ${CAMEL_MODULE_NAME}Controller.Service.GetById(ctx, ${CAMEL_MODULE_NAME}ID)

	if err != nil {
		helper.ErrorResponse(c, "Invalid request body", err, http.StatusBadRequest)
		return
	}

	responseData := Get${PASCAL_MODULE_NAME}ResponseDto{
		ID:       result.ID(),
	}

	helper.SuccessResponse(c, "Show ${CAMEL_MODULE_NAME}", responseData, http.StatusOK)
}



EOF

cat > internal/modules/$CAMEL_MODULE_NAME/application/commands/create_${CAMEL_MODULE_NAME}.go << EOF
package commands

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/dtos"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)


type Create${PASCAL_MODULE_NAME}Service struct {
	${PASCAL_MODULE_NAME}Repository   ports.${PASCAL_MODULE_NAME}Repository
	logger           logger.Logger
	tracer           trace.Tracer
}

var _ applicationPorts.${PASCAL_MODULE_NAME}Service = (*Create${PASCAL_MODULE_NAME}Service)(nil)

// NewCreate${PASCAL_MODULE_NAME}Service initializes a new Create${PASCAL_MODULE_NAME}Service with the provided ${PASCAL_MODULE_NAME}Repository.
func NewCreate${PASCAL_MODULE_NAME}Service(${CAMEL_MODULE_NAME}Repository ports.${PASCAL_MODULE_NAME}Repository, logger logger.Logger) *Create${PASCAL_MODULE_NAME}Service {
	return &Create${PASCAL_MODULE_NAME}Service{
		${PASCAL_MODULE_NAME}Repository:   ${CAMEL_MODULE_NAME}Repository,
		logger:           logger,
		tracer:           otel.Tracer("Create${PASCAL_MODULE_NAME}Service"),
	}
}

func (${CAMEL_MODULE_NAME}Service Create${PASCAL_MODULE_NAME}Service) Create${PASCAL_MODULE_NAME}(ctx context.Context, ${CAMEL_MODULE_NAME}Input dto.Create${PASCAL_MODULE_NAME}Input) error {

	ctx, span := ${CAMEL_MODULE_NAME}Service.tracer.Start(ctx, "Create${PASCAL_MODULE_NAME} - Service")
	defer span.End()

	// Create a new ${CAMEL_MODULE_NAME} using the factory
	${CAMEL_MODULE_NAME}, _ := domain.Create(
		${CAMEL_MODULE_NAME}Input.ID,
		//.. add other fields here
	)

	// todo check is the ${CAMEL_MODULE_NAME} exists in db?

	savingErr := ${CAMEL_MODULE_NAME}Service.${PASCAL_MODULE_NAME}Repository.Save(ctx, ${CAMEL_MODULE_NAME})

	if savingErr != nil {
		${CAMEL_MODULE_NAME}Service.logger.Error("Error saving ${CAMEL_MODULE_NAME}: %v", savingErr)
		return savingErr
	}

	${CAMEL_MODULE_NAME}Service.logger.Info("${PASCAL_MODULE_NAME} created successfully with ID: %s", ${CAMEL_MODULE_NAME}.ID())

	return nil
}

func (${CAMEL_MODULE_NAME}Service Create${PASCAL_MODULE_NAME}Service) GetById(ctx context.Context, id string) (*domain.${PASCAL_MODULE_NAME}, error) {

	ctx, span := ${CAMEL_MODULE_NAME}Service.tracer.Start(ctx, "GetById - Service")
	defer span.End()

	return ${CAMEL_MODULE_NAME}Service.${PASCAL_MODULE_NAME}Repository.GetByID(ctx, id)
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/entities/${CAMEL_MODULE_NAME}_entity.go << EOF
package entities


type ${PASCAL_MODULE_NAME} struct {
	ID          string    \`gorm:"primaryKey;type:uuid;default:gen_random_uuid()"\`
}

EOF

cat > internal/modules/$CAMEL_MODULE_NAME/infrastructure/persistence/gorm/mappers/${CAMEL_MODULE_NAME}_mapper.go << EOF
package mappers

import (
	domain "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
	entity "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/infrastructure/persistence/gorm/entities"
	"github.com/racibaz/go-arch/pkg/es"
)

func ToDomain(${CAMEL_MODULE_NAME}Entity entity.${PASCAL_MODULE_NAME}) domain.${PASCAL_MODULE_NAME} {


	return domain.${PASCAL_MODULE_NAME}{
		Aggregate:   es.NewAggregate(${CAMEL_MODULE_NAME}Entity.ID, domain.${PASCAL_MODULE_NAME}Aggregate),
	}

}

func ToPersistence(${CAMEL_MODULE_NAME} domain.${PASCAL_MODULE_NAME}) entity.${PASCAL_MODULE_NAME} {

	return entity.${PASCAL_MODULE_NAME}{
		ID:          ${CAMEL_MODULE_NAME}.ID(),
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

func New() *Gorm${PASCAL_MODULE_NAME}Repository {
	return &Gorm${PASCAL_MODULE_NAME}Repository{
		DB: database.Connection(),
	}
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) Save(ctx context.Context, ${CAMEL_MODULE_NAME} *domain.${PASCAL_MODULE_NAME}) error {
	var new${PASCAL_MODULE_NAME} entities.${PASCAL_MODULE_NAME}

	persistenceModel := ${CAMEL_MODULE_NAME}Mapper.ToPersistence(*${CAMEL_MODULE_NAME})

	err := repo.DB.WithContext(ctx).Create(&persistenceModel).Scan(&new${PASCAL_MODULE_NAME}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Gorm${PASCAL_MODULE_NAME}Repository) GetByID(ctx context.Context, id string) (*domain.${PASCAL_MODULE_NAME}, error) {

	var ${CAMEL_MODULE_NAME} domain.${PASCAL_MODULE_NAME}

	if err := repo.DB.WithContext(ctx).Where("id = ?", id).First(&${CAMEL_MODULE_NAME}).Error; err != nil {
		return nil, err
	}

	return &${CAMEL_MODULE_NAME}, nil
}


func (repo *Gorm${PASCAL_MODULE_NAME}Repository) IsExists(ctx context.Context, title, description string) (bool, error) {

	var ${CAMEL_MODULE_NAME} domain.${PASCAL_MODULE_NAME}

	repo.DB.WithContext(ctx).Where("title = ?", title).Where("description = ?", description).First(&${CAMEL_MODULE_NAME})

	if ${CAMEL_MODULE_NAME}.ID() != "" {
		return true, nil
	}

	return false, nil
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/domain/${CAMEL_MODULE_NAME}.go << EOF
package domain

import (
	"github.com/racibaz/go-arch/pkg/es"
)

const (
	${PASCAL_MODULE_NAME}Aggregate = "${CAMEL_MODULE_NAME}s.${PASCAL_MODULE_NAME}"
)


type ${PASCAL_MODULE_NAME} struct {
	es.Aggregate
}


// Create This factory method creates a new ${PASCAL_MODULE_NAME} with default values if you want.
func Create(id string) (*${PASCAL_MODULE_NAME}, error) {

	${CAMEL_MODULE_NAME} := &${PASCAL_MODULE_NAME}{
		Aggregate:   es.NewAggregate(id, ${PASCAL_MODULE_NAME}Aggregate),
	}

	//todo validate the ${CAMEL_MODULE_NAME} before returning it


	return ${CAMEL_MODULE_NAME}, nil
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
	IsExists(ctx context.Context, title, description string) (bool, error)
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/presentation/routes/routes.go << EOF
package routes

import (
	"github.com/gin-gonic/gin"
	${CAMEL_MODULE_NAME}Module "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}"
	${CAMEL_MODULE_NAME}Controller "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/presentation/http"
	googleGrpc "google.golang.org/grpc"

)

func Routes(router *gin.Engine) {

	module := ${CAMEL_MODULE_NAME}Module.New${PASCAL_MODULE_NAME}Module()
	new${PASCAL_MODULE_NAME}Controller := ${CAMEL_MODULE_NAME}Controller.New${PASCAL_MODULE_NAME}Controller(module.Service())

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/${CAMEL_MODULE_NAME}s")
		{
			eg.GET("/:id", new${PASCAL_MODULE_NAME}Controller.Show)
			eg.POST("/", new${PASCAL_MODULE_NAME}Controller.Store)
		}
	}
}

func GrpcRoutes(grpcServer *googleGrpc.Server) {

  // Add gRPC routes here
}

EOF


cat > internal/modules/$CAMEL_MODULE_NAME/application/ports/${CAMEL_MODULE_NAME}_service.go << EOF
package ports

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain"
)


type ${PASCAL_MODULE_NAME}Service interface {
	Create${PASCAL_MODULE_NAME}(ctx context.Context, ${CAMEL_MODULE_NAME}Dto dto.Create${PASCAL_MODULE_NAME}Input) error
	GetById(ctx context.Context, id string) (*domain.${PASCAL_MODULE_NAME}, error)
}


EOF


cat > internal/modules/$CAMEL_MODULE_NAME/application/dtos/create_${CAMEL_MODULE_NAME}_dto.go << EOF
package dto


type Create${PASCAL_MODULE_NAME}Input struct {
	ID          string // Unique identifier for the ${CAMEL_MODULE_NAME}
}

EOF



for file in controller service repository validation; do
cat > internal/modules/$CAMEL_MODULE_NAME/tests/${CAMEL_MODULE_NAME}_${file}_test.go << EOF
package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test${PASCAL_MODULE_NAME}$(echo $file | sed 's/.*/\u&/') (t *testing.T) {
	assert.True(t, true)
}
EOF
done

cat > internal/modules/$CAMEL_MODULE_NAME/module.go << EOF
package module

import (
	"github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/commands"
	${CAMEL_MODULE_NAME}Service "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/application/ports"
	${CAMEL_MODULE_NAME}Ports "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/domain/ports"
  gorm${MODULE_NAME}Repo "github.com/racibaz/go-arch/internal/modules/${CAMEL_MODULE_NAME}/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/ddd"
)

type ${PASCAL_MODULE_NAME}Module struct {
	repository ${CAMEL_MODULE_NAME}Ports.${PASCAL_MODULE_NAME}Repository
	service    ${CAMEL_MODULE_NAME}Service.${PASCAL_MODULE_NAME}Service
	log        logger.Logger
}


func New${PASCAL_MODULE_NAME}Module() *${PASCAL_MODULE_NAME}Module {

	repo := gormUserRepo.New()         // Use GORM repository for persistence
	logger, _ := logger.NewZapLogger() // Assuming NewZapLogger is a function that initializes a logger

	create${PASCAL_MODULE_NAME}Service := commands.NewCreate${PASCAL_MODULE_NAME}Service(repo, logger)

	return &${PASCAL_MODULE_NAME}Module{
		repository: repo,
		service:    create${PASCAL_MODULE_NAME}Service,
		logger:     logger,
	}
}

func (m ${PASCAL_MODULE_NAME}Module) Repository() ${CAMEL_MODULE_NAME}Ports.${PASCAL_MODULE_NAME}Repository {
	return m.repository
}

func (m ${PASCAL_MODULE_NAME}Module) Service() ${CAMEL_MODULE_NAME}Service.${PASCAL_MODULE_NAME}Service {
	return m.service
}

func (m ${PASCAL_MODULE_NAME}Module) Logger() logger.Logger {
	return m.logger
}

EOF

echo "✅ Module $PASCAL_MODULE_NAME created successfully!"
