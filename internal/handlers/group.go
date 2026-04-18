package handlers

import (
	"errors"
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupHandler struct {
	service *services.GroupService
}

var _ responses.GroupResponse

func NewGroupHandler(service *services.GroupService) *GroupHandler {
	return &GroupHandler{service: service}
}

// Create godoc
// @Summary Create a group
// @Tags group
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param group body requests.CreateGroupRequest true "Group data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/groups [post]
func (h *GroupHandler) Create(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.CreateGroupRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Create(dto, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "group created successfully",
	})
}

// Update godoc
// @Summary Update a group
// @Tags group
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param group body requests.UpdateGroupRequest true "Group data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/groups/{id} [patch]
func (h *GroupHandler) Update(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.UpdateGroupRequest
	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Update(dto, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "group not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "group updated successfully",
	})
}

// Delete godoc
// @Summary Delete a group
// @Tags group
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Group ID"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/groups/{id} [delete]
func (h *GroupHandler) Delete(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.DeleteGroupRequest
	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Delete(dto, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "group not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetGroupByID godoc
// @Summary Get group by ID
// @Tags group
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} responses.GroupResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/groups/{id} [get]
func (h *GroupHandler) GetGroupByID(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.GetGroupByID
	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	group, err := h.service.GetGroupByID(dto, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "group not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, group)
}

// GetAllGroups godoc
// @Summary List groups
// @Tags group
// @Security ApiKeyAuth
// @Produce json
// @Param name query string false "Group name filter"
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} responses.PaginatedGroupResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/groups [get]
func (h *GroupHandler) GetAllGroups(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.GetAllGroups
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	groups, err := h.service.GetAll(dto, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, groups)
}
