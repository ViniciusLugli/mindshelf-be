package responses

import (
	"time"

	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type ReceivedFriendRequestResponse struct {
	Requester UserResponse `json:"requester"`
	CreatedAt time.Time    `json:"created_at"`
}

type StatusMessageResponse struct {
	Message string `json:"message"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
}

func NewAuthResponse(token string, user models.User) AuthResponse {
	return AuthResponse{
		Token: token,
		User: UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			AvatarURL: user.AvatarURL,
		},
	}
}

func NewReceivedFriendRequestResponse(friendship models.UserFriend) ReceivedFriendRequestResponse {
	return ReceivedFriendRequestResponse{
		Requester: NewUserResponse(friendship.User),
		CreatedAt: friendship.CreatedAt,
	}
}
