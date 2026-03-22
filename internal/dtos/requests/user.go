package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetUser struct {
	ID    uuid.UUID `form:"id" binding:"uuid"`
	Email string    `form:"email"`
}

type GetAllUsers struct {
	Page  int `form:"page" binding:"min=1,required"`
	Limit int `form:"limit" binding:"min=1,max=300,required"`
}

type GetAllUsersByName struct {
	Name  string `uri:"name" binding:"required"`
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

func (d *CreateUserRequest) ToModel() models.User {
	return models.User{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}
