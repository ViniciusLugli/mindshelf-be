package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"not null"`
	Notes string `gorm:"not null"`
}
