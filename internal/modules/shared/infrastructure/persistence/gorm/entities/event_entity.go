package entities

import (
	"time"
)

type Event struct {
	StreamID      string    `gorm:"primaryKey;type:uuid;not null"`
	StreamName    string    `gorm:"type:varchar(255);not null"`
	StreamVersion string    `gorm:"type:varchar(255);not null"`
	EventID       string    `gorm:"type:varchar(255);not null"`
	EventName     string    `gorm:"type:text;not null"`
	EventData     string    `gorm:"type:text;not null"`
	OccurredAt    time.Time `gorm:"autoCreateTime:true"`
}
