package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetTask struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
}

type GetAllTasks struct {
	Title   string    `form:"title"`
	GroupID uuid.UUID `form:"group_id,parser=encoding.TextUnmarshaler" binding:"omitempty,uuid"`
	Page    int       `form:"page" binding:"min=1,required"`
	Limit   int       `form:"limit" binding:"min=1,max=300,required"`
}

type CreateTaskRequest struct {
	Title   string    `json:"title" binding:"required"`
	Notes   string    `json:"notes"`
	GroupID uuid.UUID `json:"group_id" binding:"required,uuid"`
}

type UpdateTaskRequest struct {
	ID    uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
	Title string    `json:"title"`
	Notes string    `json:"notes"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
}

func (d *CreateTaskRequest) ToModel() models.Task {
	return models.Task{
		Title:   d.Title,
		Notes:   d.Notes,
		GroupID: d.GroupID,
	}
}
