package models

import (
	"time"

	"github.com/google/uuid"
)

type MessageType string

const (
	MessageTypeText       MessageType = "text"
	MessageTypeSharedTask MessageType = "shared_task"
)

type Message struct {
	BaseModel
	Type       MessageType `gorm:"not null;default:text"`
	SenderID   uuid.UUID   `gorm:"not null;index"`
	ReceiverID uuid.UUID   `gorm:"not null;index"`
	Content    string      `gorm:"not null"`

	SharedTaskSourceID   *uuid.UUID
	SharedTaskTitle      string
	SharedTaskNotes      string
	SharedTaskGroupName  string
	SharedTaskGroupColor string

	ReceivedAt *time.Time
	ReadAt     *time.Time

	Sender   User `gorm:"foreignKey:SenderID"`
	Receiver User `gorm:"foreignKey:ReceiverID"`
}
