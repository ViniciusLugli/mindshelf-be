package responses

import "github.com/ViniciusLugli/mindshelf/internal/models"

type ChatResponse struct {
    Friend      UserResponse    `json:"friend"`
    LastMessage MessageResponse `json:"last_message"`
}

func NewChatResponse(friend models.User, last models.Message) ChatResponse {
    return ChatResponse{
        Friend:      NewUserResponse(friend),
        LastMessage: NewMessageResponse(last),
    }
}
