package queries

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

type MeQueryHandlerQuery struct {
	RefreshToken string
}

type MeQueryHandlerResponse struct {
	ID        string
	Name      string
	Email     string
	Status    domain.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
