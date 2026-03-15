package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
	DeletedAt time.Time `gorm:"index"`
}
