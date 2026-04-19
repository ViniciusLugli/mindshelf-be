package handlers

import (
	"errors"
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/gin-gonic/gin"
)

type SharedTaskHandler struct {
	service *services.MessageService
}

var _ responses.TaskResponse

func NewSharedTaskHandler(service *services.MessageService) *SharedTaskHandler {
	return &SharedTaskHandler{service: service}
}

// ImportSharedTask godoc
// @Summary Import a shared task snapshot
// @Tags shared-task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.ImportSharedTaskRequest true "Import shared task payload"
// @Success 200 {object} responses.TaskResponse
// @Success 201 {object} responses.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/shared-tasks/import [post]
func (h *SharedTaskHandler) Import(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.ImportSharedTaskRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, created, err := h.service.ImportSharedTask(userID, dto)
	if err != nil {
		if errors.Is(err, services.ErrTaskGroupNotFound) || errors.Is(err, services.ErrSharedTaskMessageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := http.StatusOK
	if created {
		status = http.StatusCreated
	}

	c.JSON(status, task)
}
