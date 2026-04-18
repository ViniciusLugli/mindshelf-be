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

type TaskHandler struct {
	service *services.TaskService
}

var _ responses.TaskResponse

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask godoc
// @Summary Create a new task
// @Tags task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param task body requests.CreateTaskRequest true "Task data"
// @Success 201 {object} responses.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tasks [post]
func (h *TaskHandler) Create(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.CreateTaskRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := h.service.Create(dto, userID)
	if err != nil {
		if errors.Is(err, services.ErrTaskGroupNotFound) {
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

	c.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Tags task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body requests.UpdateTaskRequest true "Task data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tasks/{id} [patch]
func (h *TaskHandler) Update(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.UpdateTaskRequest
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
				"error": "task not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
	})
}

// DeleteTask godoc
// @Summary Delete a task
// @Tags task
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Task ID"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tasks/{id} [delete]
func (h *TaskHandler) Delete(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.DeleteTaskRequest
	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Delete(dto, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
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

// GetTask godoc
// @Summary Get task by ID
// @Tags task
// @Security ApiKeyAuth
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} responses.TaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.GetTask
	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task, err := h.service.GetTask(dto, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "task not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetAllTasks godoc
// @Summary List tasks
// @Tags task
// @Security ApiKeyAuth
// @Produce json
// @Param title query string false "Task title filter"
// @Param group_id query string false "Group ID filter"
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} responses.PaginatedTaskResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/tasks [get]
func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var dto requests.GetAllTasks
	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tasks, err := h.service.GetAllTasks(dto, userID)
	if err != nil {
		if errors.Is(err, services.ErrTaskGroupNotFound) {
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

	c.JSON(http.StatusOK, tasks)
}
