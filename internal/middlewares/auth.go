package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ViniciusLugli/mindshelf/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "empty token",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set("userID", claims.ID)
		c.Next()
	}
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
