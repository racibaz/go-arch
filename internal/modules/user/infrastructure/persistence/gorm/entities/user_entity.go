package entities

import (
	"time"

	"github.com/racibaz/go-arch/internal/modules/user/domain"
)

type User struct {
	ID        string            `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserName  string            `gorm:"type:varchar(191);not null"`
	Email     string            `gorm:"type:varchar(191);not null"`
	Password  string            `gorm:"type:varchar(191);not null"`
	Status    domain.UserStatus `gorm:"type:int;not null"`
	CreatedAt time.Time         `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time         `gorm:"autoUpdateTime:true"`
}
