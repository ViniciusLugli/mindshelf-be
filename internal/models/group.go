package models

import "github.com/google/uuid"

type Group struct {
	BaseModel
	Name   string `gorm:"not null"`
	Color  string `gorm:"not null"`
	User   User
	UserID uuid.UUID
	Tasks  []Task
}
