package responses

import (
	"time"

	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type MessageResponse struct {
	ID         uuid.UUID    `json:"id"`
	SenderID   uuid.UUID    `json:"sender_id"`
	ReceiverID uuid.UUID    `json:"receiver_id"`
	Content    string       `json:"content"`
	CreatedAt  time.Time    `json:"created_at"`
	ReceivedAt *time.Time   `json:"received_at,omitempty"`
	ReadAt     *time.Time   `json:"read_at,omitempty"`
	Sender     UserResponse `json:"sender"`
	Receiver   UserResponse `json:"receiver"`
}

type MarkMessagesReadResponse struct {
	WithUserID uuid.UUID  `json:"with_user_id"`
	Updated    int64      `json:"updated"`
	ReadAt     *time.Time `json:"read_at,omitempty"`
}

type MessagesReadEvent struct {
	ByUserID      uuid.UUID  `json:"by_user_id"`
	WithUserID    uuid.UUID  `json:"with_user_id"`
	Updated       int64      `json:"updated"`
	ReadAt        *time.Time `json:"read_at,omitempty"`
	UpToMessageID *uuid.UUID `json:"up_to_message_id,omitempty"`
}

func NewMessageResponse(m models.Message) MessageResponse {
	return MessageResponse{
		ID:         m.ID,
		SenderID:   m.SenderID,
		ReceiverID: m.ReceiverID,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		ReceivedAt: m.ReceivedAt,
		ReadAt:     m.ReadAt,
		Sender:     NewUserResponse(m.Sender),
		Receiver:   NewUserResponse(m.Receiver),
	}
}
