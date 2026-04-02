package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetGroupByID struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

type GetAllGroupsByName struct {
	Name  string `uri:"name" binding:"required"`
	Page  int    `form:"page" binding:"min=1,required"`
	Limit int    `form:"limit" binding:"min=1,max=300,required"`
}

type GetAllGroups struct {
	Page  int `form:"page" binding:"min=1,required"`
	Limit int `form:"limit" binding:"min=1,max=300,required"`
}

type CreateGroupRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateGroupRequest struct {
	ID   uuid.UUID `json:"id" binding:"uuid,required"`
	Name string    `json:"name" binding:"required"`
}

type DeleteGroupRequest struct {
	ID uuid.UUID `form:"id" binding:"uuid,required"`
}

func (d *CreateGroupRequest) ToModel(userID uuid.UUID) models.Group {
	return models.Group{
		Name:   d.Name,
		Color:  d.Color,
		UserID: userID,
	}
}
