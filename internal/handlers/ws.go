package handlers

import (
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	util "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := strings.TrimSpace(r.Header.Get("Origin"))
		if origin == "" {
			return true
		}

		originURL, err := url.Parse(origin)
		if err != nil {
			return false
		}

		if strings.EqualFold(originURL.Host, r.Host) {
			return true
		}

		allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
		for _, allowedOrigin := range allowedOrigins {
			if strings.EqualFold(strings.TrimSpace(allowedOrigin), origin) {
				return true
			}
		}

		return false
	},
}

type WSHandler struct {
	hub    *util.Hub
	router *util.Router
}

func NewWSHandler(hub *util.Hub, router *util.Router) *WSHandler {
	return &WSHandler{hub: hub, router: router}
}

func (h *WSHandler) Handle(c *gin.Context) {
	userID, err := middlewares.GetAuthenticatedUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := util.NewClient(userID, conn, h.hub)

	h.hub.Register(client)

	go client.WritePump()
	client.ReadPump(h.router)
}
