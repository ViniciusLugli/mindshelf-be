package requests

import (
	"github.com/ViniciusLugli/mindshelf/internal/models"
	"github.com/google/uuid"
)

type GetGroupByID struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
}

type GetAllGroups struct {
	Name  string `form:"name"`
	Page  int    `form:"page" binding:"min=1,required"`
	Limit int    `form:"limit" binding:"min=1,max=300,required"`
}

type CreateGroupRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateGroupRequest struct {
	ID   uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
	Name string    `json:"name" binding:"required"`
}

type DeleteGroupRequest struct {
	ID uuid.UUID `uri:"id,parser=encoding.TextUnmarshaler" binding:"required,uuid"`
}

func (d *CreateGroupRequest) ToModel(userID uuid.UUID) models.Group {
	return models.Group{
		Name:   d.Name,
		Color:  d.Color,
		UserID: userID,
	}
}
