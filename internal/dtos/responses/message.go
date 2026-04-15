package responses

import (
	"time"

	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type MessageResponse struct {
	ID         uuid.UUID                   `json:"id"`
	Type       string                      `json:"type"`
	SenderID   uuid.UUID                   `json:"sender_id"`
	ReceiverID uuid.UUID                   `json:"receiver_id"`
	Content    string                      `json:"content"`
	CreatedAt  time.Time                   `json:"created_at"`
	ReceivedAt *time.Time                  `json:"received_at,omitempty"`
	ReadAt     *time.Time                  `json:"read_at,omitempty"`
	SharedTask *SharedTaskSnapshotResponse `json:"shared_task,omitempty"`
	Sender     UserResponse                `json:"sender"`
	Receiver   UserResponse                `json:"receiver"`
}

type SharedTaskSnapshotResponse struct {
	SourceTaskID uuid.UUID `json:"source_task_id"`
	Title        string    `json:"title"`
	Notes        string    `json:"notes"`
	GroupName    string    `json:"group_name"`
	GroupColor   string    `json:"group_color"`
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
	var sharedTask *SharedTaskSnapshotResponse
	if m.SharedTaskSourceID != nil {
		sharedTask = &SharedTaskSnapshotResponse{
			SourceTaskID: *m.SharedTaskSourceID,
			Title:        m.SharedTaskTitle,
			Notes:        m.SharedTaskNotes,
			GroupName:    m.SharedTaskGroupName,
			GroupColor:   m.SharedTaskGroupColor,
		}
	}

	return MessageResponse{
		ID:         m.ID,
		Type:       string(m.Type),
		SenderID:   m.SenderID,
		ReceiverID: m.ReceiverID,
		Content:    m.Content,
		CreatedAt:  m.CreatedAt,
		ReceivedAt: m.ReceivedAt,
		ReadAt:     m.ReadAt,
		SharedTask: sharedTask,
		Sender:     NewUserResponse(m.Sender),
		Receiver:   NewUserResponse(m.Receiver),
	}
}
