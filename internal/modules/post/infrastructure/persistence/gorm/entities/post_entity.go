package entities

import (
	"time"
)

type Post struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID      string    `gorm:"type:varchar(191);not null"`
	Title       string    `gorm:"type:varchar(191);not null"`
	Description string    `gorm:"type:varchar(191);not null"`
	Content     string    `gorm:"type:text;not null"`
	Status      int       `gorm:"type:int;not null"` // todo use value object instead of int
	CreatedAt   time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:true"`
}
