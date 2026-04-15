package util

import (
	"encoding/json"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
)

type HandlerFunc func(cl *Client, payload json.RawMessage)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) On(action string, fn HandlerFunc) {
	r.handlers[action] = fn
}

func (r *Router) Dispatch(cl *Client, msg requests.RequestMessage) {
	fn, ok := r.handlers[msg.Action]
	if !ok {
		cl.SendError(msg.Action, "unknown action")
		return
	}

	fn(cl, msg.Payload)
}
