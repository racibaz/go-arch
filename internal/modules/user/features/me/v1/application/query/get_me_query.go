package query

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

type GetMeByIdQuery struct {
	ID string
}

type GetMeByIdQueryResponse struct {
	ID        string
	UserName  string
	Email     string
	Status    domain.UserStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
