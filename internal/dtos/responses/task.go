package responses

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type TaskResponse struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Notes      string    `json:"notes"`
	GroupID    uuid.UUID `json:"group_id"`
	GroupName  string    `json:"group_name"`
	GroupColor string    `json:"group_color"`
}

func NewTaskResponse(task models.Task) TaskResponse {
	return TaskResponse{
		ID:         task.ID,
		Title:      task.Title,
		Notes:      task.Notes,
		GroupID:    task.GroupID,
		GroupName:  task.Group.Name,
		GroupColor: task.Group.Color,
	}
}
