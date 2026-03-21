package responses

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GroupResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type PaginatedGroupResponse struct {
	Data        []GroupResponse `json:"data"`
	Total       int64           `json:"total"`
	Page        int             `json:"page"`
	Limit       int             `json:"limit"`
	Total_pages int             `json:"total_pages"`
}

func NewGroupRespone(group models.Group) GroupResponse {
	return GroupResponse{
		ID:   group.ID,
		Name: group.Name,
	}
}

func NewPaginatedGroupResponse(groups []models.Group, total int64, page, limit, total_pages int) PaginatedGroupResponse {
	data := make([]GroupResponse, len(groups))
	for i, group := range groups {
		data[i] = NewGroupRespone(group)
	}

	return PaginatedGroupResponse{
		Data:        data,
		Total:       total,
		Page:        page,
		Limit:       limit,
		Total_pages: total_pages,
	}
}
