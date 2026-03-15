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

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id" binding:"required,uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email" binding:"email"`
	Password string    `json:"password"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `form:"id" binding:"required,uuid"`
}

func (d *CreateUserRequest) ToModel() models.Users {
	return models.Users{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}
