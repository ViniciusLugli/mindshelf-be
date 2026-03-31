package util

import (
	"time"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	UserID uuid.UUID
	conn   *websocket.Conn
	send   chan responses.ResponseMessage
	hub    *Hub
}

func NewClient(userID uuid.UUID, conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		ID:     uuid.New(),
		UserID: userID,
		conn:   conn,
		send:   make(chan responses.ResponseMessage, 64),
		hub:    hub,
	}
}

func (cl *Client) WritePump() {
	defer cl.conn.Close()

	for msg := range cl.send {
		if err := cl.conn.WriteJSON(msg); err != nil {
			break
		}
	}
}

func (cl *Client) ReadPump(router *Router) {
	defer func() {
		cl.hub.Unregister(cl)
		cl.conn.Close()
	}()

	cl.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	cl.conn.SetPongHandler(func(string) error {
		cl.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		var msg requests.RequestMessage
		if err := cl.conn.ReadJSON(&msg); err != nil {
			break
		}

		router.Dispatch(cl, msg)
	}
}

func (cl *Client) Send(action string, data any) {
	cl.send <- responses.ResponseMessage{Action: action, Success: true, Data: data}
}

func (cl *Client) SendError(action string, err string) {
	cl.send <- responses.ResponseMessage{Action: action, Success: false, Error: err}
}
