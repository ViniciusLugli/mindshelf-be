package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetUserByID struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
}

type GetAllUsers struct {
	Name  string `form:"name"`
	Page  int    `form:"page" binding:"min=1,required"`
	Limit int    `form:"limit" binding:"min=1,max=300,required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type FriendRequest struct {
	FriendID uuid.UUID `json:"friend_id" binding:"required,uuid"`
}

func (d *CreateUserRequest) ToModel() models.User {
	return models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}
