package models

type Task struct {
	BaseModel
	Title string `gorm:"not null"`
	Notes string `gorm:"not null"`
}
