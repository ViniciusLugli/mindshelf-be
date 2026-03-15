package models

type Users struct {
	BaseModel
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
	Groups   []Group
}
