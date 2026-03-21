package responses

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type TaskResponse struct {
	Title   string    `json:"title"`
	Notes   string    `json:"notes"`
	GroupID uuid.UUID `json:"group_id"`
}

func NewTaskResponse(task models.Task) TaskResponse {
	return TaskResponse{
		Title:   task.Title,
		Notes:   task.Notes,
		GroupID: task.GroupID,
	}
}
