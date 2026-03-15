package models

import "github.com/google/uuid"

type Task struct {
	BaseModel
	Title   string `gorm:"not null"`
	Notes   string `gorm:"not null"`
	Group   Group
	GroupID uuid.UUID
}
