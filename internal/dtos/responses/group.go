package responses

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GroupResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewGroupRespone(group models.Group) GroupResponse {
	return GroupResponse{
		ID:   group.ID,
		Name: group.Name,
	}
}
