package commands

import (
	"encoding/json"
	"github.com/racibaz/go-arch/internal/modules/post/application/dtos"
	applicationPorts "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/domain"
	"github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/racibaz/go-arch/pkg/messaging"
	"github.com/racibaz/go-arch/pkg/messaging/rabbitmq"
	"log"
	"time"
)

type CreatePostService struct {
	PostRepository   ports.PostRepository
	logger           logger.Logger
	messagePublisher rabbitmq.MessagePublisher
}

var _ applicationPorts.PostService = (*CreatePostService)(nil)

// NewCreatePostService initializes a new CreatePostService with the provided PostRepository.
func NewCreatePostService(postRepository ports.PostRepository, logger logger.Logger, messagePublisher rabbitmq.MessagePublisher) *CreatePostService {
	return &CreatePostService{
		PostRepository:   postRepository,
		logger:           logger,
		messagePublisher: messagePublisher,
	}
}

func (postService CreatePostService) CreatePost(postInput dto.CreatePostInput) error {

	// Create a new post using the factory
	post, _ := domain.Create(
		postInput.ID,
		postInput.UserID,
		postInput.Title,
		postInput.Description,
		postInput.Content,
		postInput.Status,
		time.Now(),
		time.Now(),
	)

	// check is the post exists in db?
	isExists, err := postService.PostRepository.IsExists(post.Title, post.Description)

	if err != nil {
		return err
	}

	// If the post already exists, return an error
	if isExists {
		postService.logger.Info("Post already exists with title: %s and description: %s", post.Title, post.Description)
		return domain.ErrAlreadyExists
	}

	savingErr := postService.PostRepository.Save(post)

	if savingErr != nil {
		postService.logger.Error("Error saving post: %v", savingErr)
		return savingErr
	}

	payload, err := json.Marshal(post)
	if err != nil {
		log.Printf("Error marshalling payload: %v", err)
		return err
	}

	postService.messagePublisher.PublishEvent("", messaging.MessagePayload{
		OwnerID: post.UserID, // todo it should get auth user id
		Data:    payload,
	})

	postService.logger.Info("Post created successfully with ID: %s", post.ID())

	return nil
}

func (postService CreatePostService) GetById(id string) (*domain.Post, error) {

	return postService.PostRepository.GetByID(id)
}
