package entities

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/post/domain"
)

type Post struct {
	ID          string            `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string            `gorm:"type:varchar(191);not null"`
	Title       string            `gorm:"type:varchar(191);not null"`
	Description string            `gorm:"type:varchar(191);not null"`
	Content     string            `gorm:"type:text;not null"`
	Status      domain.PostStatus `gorm:"type:int;not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime:true"`
}
