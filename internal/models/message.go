package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	BaseModel
	SenderID   uuid.UUID `gorm:"not null;index"`
	ReceiverID uuid.UUID `gorm:"not null;index"`
	Content    string    `gorm:"not null"`
	ReceivedAt *time.Time
	ReadAt     *time.Time

	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}
