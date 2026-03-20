package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetGroup struct {
	Name string    `form:"name"`
	ID   uuid.UUID `form:"id"`
}

type GetAllGroups struct {
	Page  int `form:"page" binding:"min=1,required"`
	Limit int `form:"limit" binding:"min=1,max=300,required"`
}

type CreateGroupRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"uuid,required"`
	Name   string    `json:"name" binding:"required"`
}

type UpdateGroupRequest struct {
	ID   uuid.UUID `json:"id" binding:"uuid,required"`
	Name string    `json:"name" binding:"required"`
}

type DeleteGroupRequest struct {
	ID uuid.UUID `form:"id" binding:"uuid,required"`
}

func (d *CreateGroupRequest) ToModel() models.Group {
	return models.Group{
		Name:   d.Name,
		UserID: d.UserID,
	}
}
