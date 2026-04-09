package middlewares

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ViniciusLugli/mindshelf/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var authCookieNames = []string{"token", "auth_token", "access_token", "jwt"}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := extractToken(c)
		if err != nil {
			slog.Warn(
				"unauthorized request",
				"request_id", GetRequestID(c),
				"reason", err.Error(),
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "empty token",
			})
			return
		}

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			slog.Warn(
				"unauthorized request",
				"request_id", GetRequestID(c),
				"reason", "invalid token",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}

func extractToken(c *gin.Context) (string, error) {
	if token := normalizeToken(c.GetHeader("Authorization")); token != "" {
		return token, nil
	}

	for _, cookieName := range authCookieNames {
		cookieValue, err := c.Cookie(cookieName)
		if err != nil {
			continue
		}

		if token := normalizeToken(cookieValue); token != "" {
			return token, nil
		}
	}

	return "", errors.New("missing authorization token")
}

func normalizeToken(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}

	if strings.HasPrefix(strings.ToLower(raw), "bearer ") {
		return strings.TrimSpace(raw[7:])
	}

	return raw
}

func GetAuthenticatedUserID(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return uuid.Nil, errors.New("user not authenticated")
	}

	id, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("userID invalid in context")
	}

	return id, nil
}
