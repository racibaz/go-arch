package post_module

import (
	postService "github.com/racibaz/go-arch/internal/modules/post/application/ports"
	"github.com/racibaz/go-arch/internal/modules/post/application/usecases"
	postRepository "github.com/racibaz/go-arch/internal/modules/post/domain/ports"
	gromPostRepo "github.com/racibaz/go-arch/internal/modules/post/infrastructure/persistence/gorm/repositories"
)

type PostModule struct {
	repository postRepository.PostRepository
	service    postService.PostService
}

func NewPostModule() *PostModule {
	//repo := in_memory.New()
	repo := gromPostRepo.New() // Use GORM repository for persistence
	service := usecases.NewCreatetPostUseCase(repo)

	return &PostModule{
		repository: repo,
		service:    service,
	}
}

func (m *PostModule) GetRepository() postRepository.PostRepository {
	return m.repository
}

func (m *PostModule) GetService() postService.PostService {
	return m.service
}
