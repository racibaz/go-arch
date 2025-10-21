package grpc

import (
	"context"
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	proto "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"github.com/racibaz/go-arch/pkg/uuid"
	"google.golang.org/grpc"
	"time"
)

type PostGrpcController struct {
	Service ports.PostService //todo Service name will be changed to PostService and it should be camelCase
	proto.UnimplementedPostServiceServer
}

func NewPostGrpcController(grpc *grpc.Server, postService ports.PostService) {
	gRPController := &PostGrpcController{
		Service: postService,
	}

	// Register the gRPC service with the server
	proto.RegisterPostServiceServer(grpc, gRPController)
}

func (controller *PostGrpcController) CreatePost(ctx context.Context, in *proto.Post) (*proto.CreatePostResponse, error) {

	postId := uuid.NewID()

	input := dto.CreatePostInput{
		ID:          postId, //todo value object olmalı
		Title:       in.Title,
		Description: in.Description,
		Content:     in.Content,
		Status:      postValueObject.PostStatus(postValueObject.PostStatusDraft),
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err := controller.Service.CreatePost(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.CreatePostResponse{
		Id: postId,
	}, nil
}
