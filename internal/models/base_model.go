package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
	DeletedAt *time.Time `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
