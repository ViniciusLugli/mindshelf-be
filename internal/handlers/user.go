package handlers

import (
	"errors"
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

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
