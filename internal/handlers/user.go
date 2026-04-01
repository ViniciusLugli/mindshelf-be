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

type UserHandler struct {
	service *services.UserService
}

var _ responses.UserResponse

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Update godoc
// @Summary Update authenticated user
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body requests.UpdateUserRequest true "User data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/update [patch]
func (h *UserHandler) Update(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	var dto requests.UpdateUserRequest
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.Update(dto, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated succesfully",
	})
}

// Delete godoc
// @Summary Delete authenticated user
// @Tags user
// @Security ApiKeyAuth
// @Produce json
// @Success 204 {string} string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/delete [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	if err := h.service.Delete(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetUser godoc
// @Summary Get user by id or email
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id query string false "User ID"
// @Param email query string false "User email"
// @Success 200 {object} responses.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/ [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	var dto requests.GetUser
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.GetUser(dto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})

			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})

			return
		}
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get paginated users
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} responses.PaginatedUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/all [get]

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var dto requests.GetAllUsers
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := h.service.GetAllUsers(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

// GetAllUsersByName godoc
// @Summary Get paginated users by name
// @Tags user
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} responses.PaginatedUserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/{name} [get]
func (h *UserHandler) GetAllUsersByName(c *gin.Context) {
	var dto requests.GetAllUsersByName

	if err := c.ShouldBindUri(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindQuery(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.service.GetAllUsersByName(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
