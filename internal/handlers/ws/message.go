package wsHandler

import (
	"encoding/json"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	util "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
	ws "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
)

type ChatHandler struct {
	msgService *services.MessageService
	hub        *util.Hub
}

func NewChatHandler(msgSvc *services.MessageService, hub *util.Hub) *ChatHandler {
	return &ChatHandler{msgService: msgSvc, hub: hub}
}

func (h *ChatHandler) Register(r *ws.Router) {
	r.On("send_message", h.SendMessage)
	r.On("get_conversation", h.GetConversation)
	r.On("get_chats", h.GetChats)
}

func (h *ChatHandler) SendMessage(cl *ws.Client, payload json.RawMessage) {
	var dto requests.SendChatRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("send_message", "invalid payload")
		return
	}

	msg, err := h.msgService.SendMessage(cl.UserID, dto)
	if err != nil {
		cl.SendError("send_message", err.Error())
		return
	}

	cl.Send("message_sent", msg)

	h.hub.SendToUser(dto.ToUserID, "message_received", msg)
}

// Buscar histórico de conversa com um usuário
func (h *ChatHandler) GetConversation(cl *ws.Client, payload json.RawMessage) {
	var dto requests.GetChatRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("get_conversation", "invalid payload")
		return
	}

	messages, err := h.msgService.GetConversation(cl.UserID, dto.WithUserID)
	if err != nil {
		cl.SendError("get_conversation", err.Error())
		return
	}

	cl.Send("get_conversation", messages)
}

func (h *ChatHandler) GetChats(cl *ws.Client, _ json.RawMessage) {
	chats, err := h.msgService.GetChats(cl.UserID)
	if err != nil {
		cl.SendError("get_chats", err.Error())
		return
	}

	cl.Send("get_chats", chats)
}
