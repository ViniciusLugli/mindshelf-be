package wsHandler

import (
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/ViniciusLugli/mindshelf/internal/dtos/responses"
	"github.com/gin-gonic/gin"
)

var _ requests.SendChatRequest
var _ requests.ShareTaskRequest
var _ requests.RequestMessage
var _ responses.ResponseMessage
var _ responses.MessageResponse
var _ responses.ChatResponse
var _ responses.MarkMessagesReadResponse
var _ responses.ReceivedFriendRequestResponse
var _ responses.StatusMessageResponse

func RegisterWebsocketDocs(r *gin.Engine) {
	r.GET("/ws/connect", ConnectWebSocketDoc)
	r.POST("/ws/send_message", SendMessageDoc)
	r.POST("/ws/share_task", ShareTaskDoc)
	r.POST("/ws/get_conversation", GetConversationDoc)
	r.GET("/ws/get_chats", GetChatsDoc)
	r.POST("/ws/mark_messages_read", MarkMessagesReadDoc)

	r.POST("/ws/send_friend_request", SendFriendRequestDoc)
	r.POST("/ws/accept_friend_request", AcceptFriendRequestDoc)
	r.POST("/ws/reject_friend_request", RejectFriendRequestDoc)
	r.POST("/ws/remove_friend", RemoveFriendDoc)
	r.GET("/ws/get_friends", GetFriendsDoc)
	r.GET("/ws/get_pending_friend_requests", GetPendingFriendRequestsDoc)
}

// ConnectWebSocketDoc godoc
// @Summary Connect to the authenticated WebSocket
// @Description Open a WebSocket connection on `/api/ws` using the same authentication used by protected HTTP routes. After connecting, every client message must use the envelope `{"action":"event_name","payload":{...}}`. Every server response uses the envelope `{"action":"event_name","success":true|false,"data":...,"error":"..."}`.
// @Tags websocket
// @Security ApiKeyAuth
// @Produce json
// @Success 101 {string} string "Switching Protocols"
// @Router /api/ws [get]
func ConnectWebSocketDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Open a websocket connection on '/api/ws'."})
}

// SendMessageDoc godoc
// @Summary WebSocket event: send_message
// @Description Send `{"action":"send_message","payload":{...}}` after connecting to `/api/ws`. The sender receives the `message_sent` event and the recipient receives the `message_received` event. Both events carry a `responses.MessageResponse` in `data`.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.SendChatRequest true "Send message payload"
// @Success 200 {object} responses.MessageResponse
// @Router /ws/send_message [post]
func SendMessageDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'send_message' event."})
}

// ShareTaskDoc godoc
// @Summary WebSocket event: share_task
// @Description Send `{"action":"share_task","payload":{...}}` after connecting to `/api/ws`. The task must belong to the authenticated sender. The server persists a chat message with `type = shared_task`, returns `message_sent` to the sender, and broadcasts `message_received` to the recipient. The payload in `data` is `responses.MessageResponse` with the `shared_task` snapshot filled.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.ShareTaskRequest true "Share task payload"
// @Success 200 {object} responses.MessageResponse
// @Router /ws/share_task [post]
func ShareTaskDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'share_task' event."})
}

// GetConversationDoc godoc
// @Summary WebSocket event: get_conversation
// @Description Send `{"action":"get_conversation","payload":{"with_user_id":"..."}}` after connecting to `/api/ws`. The server responds to the same client with the `get_conversation` event. The event `data` contains the full ordered conversation as `[]responses.MessageResponse`.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.GetChatRequest true "Get conversation payload"
// @Success 200 {array} responses.MessageResponse
// @Router /ws/get_conversation [post]
func GetConversationDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_conversation' event."})
}

// GetChatsDoc godoc
// @Summary WebSocket event: get_chats
// @Description Send `{"action":"get_chats","payload":null}` after connecting to `/api/ws`. The server responds to the same client with the `get_chats` event. The event `data` contains `[]responses.ChatResponse` with friend info, last message, and unread count.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} responses.ChatResponse
// @Router /ws/get_chats [get]
func GetChatsDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_chats' event."})
}

// MarkMessagesReadDoc godoc
// @Summary WebSocket event: mark_messages_read
// @Description Send `{"action":"mark_messages_read","payload":{...}}` after connecting to `/api/ws`. The sender receives `mark_messages_read` with `responses.MarkMessagesReadResponse`. If at least one message changes, the other participant also receives the `messages_read` broadcast.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.MarkMessagesReadRequest true "Mark messages as read payload"
// @Success 200 {object} responses.MarkMessagesReadResponse
// @Router /ws/mark_messages_read [post]
func MarkMessagesReadDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'mark_messages_read' event."})
}

// SendFriendRequestDoc godoc
// @Summary WebSocket event: send_friend_request
// @Description Send `{"action":"send_friend_request","payload":{"friend_id":"..."}}` after connecting to `/api/ws`. The sender receives `send_friend_request` with a simple confirmation message when the pending invitation is created.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} responses.StatusMessageResponse
// @Router /ws/send_friend_request [post]
func SendFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'send_friend_request' event."})
}

// AcceptFriendRequestDoc godoc
// @Summary Accept friend request over WebSocket
// @Description Send `{"action":"accept_friend_request","payload":{"friend_id":"..."}}` after connecting to `/api/ws`. The authenticated user accepts a previously received pending invitation from `friend_id`. The server replies with `accept_friend_request` and a confirmation message.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} responses.StatusMessageResponse
// @Router /ws/accept_friend_request [post]
func AcceptFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'accept_friend_request' event."})
}

// RejectFriendRequestDoc godoc
// @Summary Reject friend request over WebSocket
// @Description Send `{"action":"reject_friend_request","payload":{"friend_id":"..."}}` after connecting to `/api/ws`. The authenticated user rejects a previously received pending invitation from `friend_id`. The server replies with `reject_friend_request` and a confirmation message.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} responses.StatusMessageResponse
// @Router /ws/reject_friend_request [post]
func RejectFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'reject_friend_request' event."})
}

// RemoveFriendDoc godoc
// @Summary Remove a friend over WebSocket
// @Description Send `{"action":"remove_friend","payload":{"friend_id":"..."}}` after connecting to `/api/ws`. The server removes the friendship in either direction and replies with `remove_friend` and a confirmation message.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Remove friend payload"
// @Success 200 {object} responses.StatusMessageResponse
// @Router /ws/remove_friend [post]
func RemoveFriendDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'remove_friend' event."})
}

// GetFriendsDoc godoc
// @Summary Get friends list over WebSocket
// @Description Send `{"action":"get_friends","payload":null}` after connecting to `/api/ws`. The server responds with `get_friends` and `[]responses.UserResponse` in `data`.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} responses.UserResponse
// @Router /ws/get_friends [get]
func GetFriendsDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_friends' event."})
}

// GetPendingFriendRequestsDoc godoc
// @Summary WebSocket event: get_pending_friend_requests
// @Description Send `{"action":"get_pending_friend_requests","payload":null}` after connecting to `/api/ws`. The server responds with `get_pending_friend_requests` and `[]responses.ReceivedFriendRequestResponse` in `data`, listing only pending invitations received by the authenticated user.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {array} responses.ReceivedFriendRequestResponse
// @Router /ws/get_pending_friend_requests [get]
func GetPendingFriendRequestsDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_pending_friend_requests' event."})
}
