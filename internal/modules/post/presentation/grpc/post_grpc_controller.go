package grpc

import (
	"context"
	"fmt"
	"time"

	dto "github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	proto "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type PostGrpcController struct {
	Service ports.PostService // todo Service name will be changed to PostService and it should be camelCase
	proto.UnimplementedPostServiceServer
}

func NewPostGrpcController(grpc *grpc.Server, postService ports.PostService) {
	gRPController := &PostGrpcController{
		Service: postService,
	}

	// Register the gRPC service with the server
	proto.RegisterPostServiceServer(grpc, gRPController)
}

func (controller *PostGrpcController) CreatePost(
	ctx context.Context,
	in *proto.CreatePostInput,
) (*proto.CreatePostResponse, error) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"PostModule - gRPC - %s - %s",
		helper.StructName(controller),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(ctx, path)
	defer span.End()

	newUuid := uuid.NewID()

	input := dto.CreatePostInput{
		ID:          newUuid,   // todo it should be value object
		UserID:      in.UserID, // todo it should be value object
		Title:       in.Title,
		Description: in.Description,
		Content:     in.Content,
		Status:      postValueObject.PostStatusDraft,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err := controller.Service.CreatePost(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.CreatePostResponse{
		Id: newUuid,
	}, nil
}
