package wsHandler

import (
	"encoding/json"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	ws "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
)

type FriendHandlers struct {
	service *services.UserService
}

func NewFriendHandlers(s *services.UserService) *FriendHandlers {
	return &FriendHandlers{service: s}
}

func (h *FriendHandlers) Register(r *ws.Router) {
	r.On("send_friend_request", h.SendFriendRequest)
	r.On("accept_friend_request", h.AcceptFriendRequest)
	r.On("reject_friend_request", h.RejectFriendRequest)
	r.On("remove_friend", h.RemoveFriend)
	r.On("get_friends", h.GetFriends)
	r.On("get_pending_friend_requests", h.GetPendingFriendRequests)
}

func (h *FriendHandlers) SendFriendRequest(cl *ws.Client, payload json.RawMessage) {
	var dto requests.FriendRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("send_friend_request", "invalid payload")
		return
	}

	err := h.service.SendFriendRequest(cl.UserID, dto)
	if err != nil {
		cl.SendError("send_friend_request", err.Error())
		return
	}

	cl.Send("send_friend_request", "request sent")
}

func (h *FriendHandlers) AcceptFriendRequest(cl *ws.Client, payload json.RawMessage) {
	var dto requests.FriendRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("accept_friend_request", "invalid payload")
		return
	}

	err := h.service.AcceptFriendRequest(cl.UserID, dto)
	if err != nil {
		cl.SendError("accept_friend_request", err.Error())
		return
	}

	cl.Send("accept_friend_request", "friendship created")
}

func (h *FriendHandlers) RejectFriendRequest(cl *ws.Client, payload json.RawMessage) {
	var dto requests.FriendRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("reject_friend_request", "invalid payload")
		return
	}

	err := h.service.RejectFriendRequest(cl.UserID, dto)
	if err != nil {
		cl.SendError("reject_friend_request", err.Error())
		return
	}

	cl.Send("reject_friend_request", "friendship rejected")
}

func (h *FriendHandlers) RemoveFriend(cl *ws.Client, payload json.RawMessage) {
	var dto requests.FriendRequest
	if err := json.Unmarshal(payload, &dto); err != nil {
		cl.SendError("remove_friend", "invalid payload")
		return
	}

	err := h.service.RemoveFriend(cl.UserID, dto)
	if err != nil {
		cl.SendError("remove_friend", err.Error())
		return
	}

	cl.Send("remove_friend", "friend removed")
}

func (h *FriendHandlers) GetFriends(cl *ws.Client, payload json.RawMessage) {
	friends, err := h.service.GetFriends(cl.UserID)
	if err != nil {
		cl.SendError("get_friends", err.Error())
		return
	}

	cl.Send("get_friends", friends)
}

func (h *FriendHandlers) GetPendingFriendRequests(cl *ws.Client, payload json.RawMessage) {
	friendRequests, err := h.service.GetPendingFriendRequests(cl.UserID)
	if err != nil {
		cl.SendError("get_pending_friend_requests", err.Error())
		return
	}

	cl.Send("get_pending_friend_requests", friendRequests)
}
