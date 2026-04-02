package wsHandler

import (
	"encoding/json"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
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
	r.On("mark_messages_read", h.MarkMessagesRead)
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

func (h *ChatHandler) MarkMessagesRead(cl *ws.Client, payload json.RawMessage) {
	var dto requests.MarkMessagesReadRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("mark_messages_read", "invalid payload")
		return
	}

	result, err := h.msgService.MarkMessagesAsRead(cl.UserID, dto)
	if err != nil {
		cl.SendError("mark_messages_read", err.Error())
		return
	}

	cl.Send("mark_messages_read", result)

	if result.Updated > 0 {
		h.hub.SendToUser(dto.WithUserID, "messages_read", responses.MessagesReadEvent{
			ByUserID:      cl.UserID,
			WithUserID:    dto.WithUserID,
			Updated:       result.Updated,
			ReadAt:        result.ReadAt,
			UpToMessageID: dto.UpToMessageID,
		})
	}
}
