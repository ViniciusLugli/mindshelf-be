package models

type Group struct {
	BaseModel
	Name  string `gorm:"not null"`
	Tasks []Task
}
