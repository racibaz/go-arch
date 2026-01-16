package grpc

import (
	"context"
	"fmt"
	"time"

	postValueObject "github.com/racibaz/go-arch/internal/modules/post/domain"
	proto2 "github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/adapters/endpoints/grpc/proto"
	"github.com/racibaz/go-arch/internal/modules/post/features/creatingpost/v1/application/commands"
	"github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/pkg/config"
	"github.com/racibaz/go-arch/pkg/helper"
	"github.com/racibaz/go-arch/pkg/uuid"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

type CreatePostHandler struct {
	Handler ports.CommandHandler[commands.CreatePostCommandV1]
	proto2.UnimplementedPostServiceServer
}

func NewCreatePostHandler(
	grpc *grpc.Server,
	postHandler ports.CommandHandler[commands.CreatePostCommandV1],
) {
	handler := &CreatePostHandler{
		Handler: postHandler,
	}

	// Register the gRPC handler with the server
	proto2.RegisterPostServiceServer(grpc, handler)
}

func (h *CreatePostHandler) CreatePost(
	ctx context.Context,
	in *proto2.CreatePostInput,
) (*proto2.CreatePostResponse, error) {
	tracer := otel.Tracer(config.Get().App.Name)
	path := fmt.Sprintf(
		"PostModule - gRPC - %s - %s",
		helper.StructName(h),
		helper.CurrentFuncName(),
	)
	ctx, span := tracer.Start(ctx, path)
	defer span.End()

	newUuid := uuid.NewID()

	input := commands.CreatePostCommandV1{
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

	return &proto2.CreatePostResponse{
		Id: newUuid,
	}, nil
}
