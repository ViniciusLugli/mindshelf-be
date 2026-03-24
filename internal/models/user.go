package models

type User struct {
	BaseModel
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	AvatarURL string
	Groups    []Group
}
