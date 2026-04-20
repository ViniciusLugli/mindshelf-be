package handlers

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/gin-gonic/gin"
)

const authCookieName = "mindshelf_token"
const authCookieMaxAgeSeconds = 30 * 24 * 60 * 60

type AuthHandler struct {
	service *services.AuthService
}

var _ responses.AuthResponse

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequest true "User data"
// @Success 201 {object} responses.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var dto requests.CreateUserRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Register(dto)
	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyInUse) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	setAuthCookie(c, user.Token)

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary Login and receive auth token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body requests.LoginRequest true "Credentials"
// @Success 200 {object} responses.AuthResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var dto requests.LoginRequest
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Login(dto)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	setAuthCookie(c, user.Token)

	c.JSON(http.StatusOK, user)
}

func setAuthCookie(c *gin.Context, token string) {
	isSecure := c.Request.TLS != nil || strings.EqualFold(c.GetHeader("X-Forwarded-Proto"), "https") || os.Getenv("COOKIE_SECURE") == "true"
	domain := strings.TrimSpace(os.Getenv("COOKIE_DOMAIN"))

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(authCookieName, token, authCookieMaxAgeSeconds, "/", domain, isSecure, true)
}
