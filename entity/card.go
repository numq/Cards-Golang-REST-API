package entity

import (
	"time"

	"github.com/google/uuid"
)

type Card struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	FrontText string    `gorm:"size:255;not null;"`
	BackText  string    `gorm:"size:255;not null;"`
	CreatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
