package responses

import "github.com/ViniciusLugli/mindshelf/internal/models"

type ChatResponse struct {
	Friend      UserResponse    `json:"friend"`
	LastMessage MessageResponse `json:"last_message"`
	UnreadCount int64           `json:"unread_count"`
}

func NewChatResponse(friend models.User, last models.Message, unreadCount int64) ChatResponse {
	return ChatResponse{
		Friend:      NewUserResponse(friend),
		LastMessage: NewMessageResponse(last),
		UnreadCount: unreadCount,
	}
}
