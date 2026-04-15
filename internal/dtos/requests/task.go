package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetTask struct {
	ID uuid.UUID `form:"id" binding:"required,uuid"`
}

type GetAllTasks struct {
	Page  int `form:"page" binding:"min=1,required"`
	Limit int `form:"limit" binding:"min=1,max=300,required"`
}

type GetAllTasksByTitle struct {
	Title string `uri:"title" binding:"required"`
	Page  int    `form:"page" binding:"min=1,required"`
	Limit int    `form:"limit" binding:"min=1,max=300,required"`
}

type GetAllTasksByGroup struct {
	GroupID uuid.UUID `uri:"groupID" binding:"required,uuid"`
	Page    int       `form:"page" binding:"min=1,required"`
	Limit   int       `form:"limit" binding:"min=1,max=300,required"`
}

type CreateTaskRequest struct {
	Title   string    `json:"title" binding:"required"`
	Notes   string    `json:"notes"`
	GroupID uuid.UUID `json:"group_id" binding:"required,uuid"`
}

type UpdateTaskRequest struct {
	ID    uuid.UUID `json:"id" binding:"required,uuid"`
	Title string    `json:"title"`
	Notes string    `json:"notes"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

func (d *CreateTaskRequest) ToModel() models.Task {
	return models.Task{
		Title:   d.Title,
		Notes:   d.Notes,
		GroupID: d.GroupID,
	}
}
