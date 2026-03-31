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
