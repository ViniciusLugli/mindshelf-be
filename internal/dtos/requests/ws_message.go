package requests

import (
	"encoding/json"

	"github.com/google/uuid"
)

type RequestMessage struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

type GetChatRequest struct {
	WithUserID uuid.UUID `json:"with_user_id"`
}

type SendChatRequest struct {
	ToUserID uuid.UUID `json:"to_user_id"`
	Content  string    `json:"content"`
}

type MarkMessagesReadRequest struct {
	WithUserID    uuid.UUID  `json:"with_user_id"`
	UpToMessageID *uuid.UUID `json:"up_to_message_id,omitempty"`
}
