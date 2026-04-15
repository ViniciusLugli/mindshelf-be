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

const authHeaderName = "Authorization"
const authCookieName = "mindshelf_token"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := extractToken(c)
		if err != nil {
			hasCookieHeader := strings.TrimSpace(c.GetHeader("Cookie")) != ""
			hasAuthorizationHeader := strings.TrimSpace(c.GetHeader(authHeaderName)) != ""

			slog.Warn(
				"unauthorized request",
				"request_id", GetRequestID(c),
				"reason", err.Error(),
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"has_cookie_header", hasCookieHeader,
				"has_authorization_header", hasAuthorizationHeader,
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
	if token := normalizeToken(c.GetHeader(authHeaderName)); token != "" {
		return token, nil
	}

	cookieValue, err := c.Cookie(authCookieName)
	if err == nil {
		if token := normalizeToken(cookieValue); token != "" {
			return token, nil
		}
	}

	if token := tokenFromRawCookieHeader(c.GetHeader("Cookie"), authCookieName); token != "" {
		return token, nil
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

func tokenFromRawCookieHeader(rawHeader, cookieName string) string {
	if strings.TrimSpace(rawHeader) == "" {
		return ""
	}

	cookies := strings.Split(rawHeader, ";")
	for _, cookie := range cookies {
		parts := strings.SplitN(strings.TrimSpace(cookie), "=", 2)
		if len(parts) != 2 {
			continue
		}

		if strings.TrimSpace(parts[0]) != cookieName {
			continue
		}

		return normalizeToken(parts[1])
	}

	return ""
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
