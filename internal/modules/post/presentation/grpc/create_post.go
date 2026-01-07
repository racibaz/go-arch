package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/racibaz/go-arch/internal/modules/post/application/command"
	"github.com/racibaz/go-arch/internal/modules/post/application/ports"
	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	proto "github.com/racibaz/go-arch/internal/modules/post/presentation/grpc/proto"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type CreatePostHandler struct {
	Handler ports.CommandHandler[command.CreatePostCommand]
	proto.UnimplementedPostServiceServer
}

func NewCreatePostHandler(
	grpc *grpc.Server,
	postHandler ports.CommandHandler[command.CreatePostCommand],
) {
	handler := &CreatePostHandler{
		Handler: postHandler,
	}

	// Register the gRPC handler with the server
	proto.RegisterPostServiceServer(grpc, handler)
}

func (h *CreatePostHandler) CreatePost(
	ctx context.Context,
	in *proto.CreatePostInput,
) (*proto.CreatePostResponse, error) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"PostModule - gRPC - %s - %s",
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(ctx, path)
	defer span.End()

	newUuid := uuid.NewID()

	input := command.CreatePostCommand{
		ID:          newUuid,   // todo it should be value object
		UserID:      in.UserID, // todo it should be value object
		Title:       in.Title,
		Description: in.Description,
		Content:     in.Content,
		Status:      postValueObject.PostStatusDraft,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err := h.Handler.Handle(ctx, input)
	if err != nil {
		return nil, err
	}

	return &proto.CreatePostResponse{
		Id: newUuid,
	}, nil
}
