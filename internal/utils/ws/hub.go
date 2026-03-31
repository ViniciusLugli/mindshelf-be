package util

import (
	"sync"

	"github.com/google/uuid"
)

type Hub struct {
	clients map[uuid.UUID]*Client
	mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{clients: make(map[uuid.UUID]*Client)}
}

func (h *Hub) Register(cl *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[cl.UserID] = cl
}

func (h *Hub) Unregister(cl *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.clients[cl.UserID]; ok {
		delete(h.clients, cl.UserID)
		close(cl.send)
	}
}

func (h *Hub) SendToUser(userID uuid.UUID, action string, data any) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if cl, ok := h.clients[userID]; ok {
		cl.Send(action, data)
	}
}
