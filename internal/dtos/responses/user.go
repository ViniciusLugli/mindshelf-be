package responses

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type PaginatedUserResponse struct {
	Data        []UserResponse `json:"data"`
	Total       int64          `json:"total"`
	Page        int            `json:"page"`
	Limit       int            `json:"limit"`
	Total_pages int            `json:"total_pages"`
}

func NewUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewPaginatedUserResponse(users []models.User, total int64, page, limit, total_pages int) PaginatedUserResponse {
	data := make([]UserResponse, len(users))
	for i, user := range users {
		data[i] = NewUserResponse(user)
	}

	return PaginatedUserResponse{
		Data:        data,
		Total:       total,
		Page:        page,
		Limit:       limit,
		Total_pages: total_pages,
	}
}
