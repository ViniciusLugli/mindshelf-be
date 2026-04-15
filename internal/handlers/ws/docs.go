package wsHandler

import (
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/dtos/requests"
	"github.com/gin-gonic/gin"
)

var _ requests.SendChatRequest
var _ requests.ShareTaskRequest

func RegisterWebsocketDocs(r *gin.Engine) {
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
}

// SendMessageDoc godoc
// @Summary Send a chat message over WebSocket
// @Description Connect to the WebSocket, send an event named `send_message` with this payload. Server will broadcast `message_received` to recipient.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.SendChatRequest true "Send message payload"
// @Success 200 {object} map[string]string
// @Router /ws/send_message [post]
func SendMessageDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'send_message' event."})
}

// ShareTaskDoc godoc
// @Summary Share a task over WebSocket
// @Description Connect to the WebSocket, send an event named `share_task` with a task snapshot payload. Server will persist the shared task in chat history and broadcast it to the recipient.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.ShareTaskRequest true "Share task payload"
// @Success 200 {object} map[string]string
// @Router /ws/share_task [post]
func ShareTaskDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'share_task' event."})
}

// GetConversationDoc godoc
// @Summary Get conversation messages over WebSocket
// @Description Send event `get_conversation` with `with_user_id` to receive conversation history.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.GetChatRequest true "Get conversation payload"
// @Success 200 {object} map[string]string
// @Router /ws/get_conversation [post]
func GetConversationDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_conversation' event."})
}

// GetChatsDoc godoc
// @Summary Get chats list over WebSocket
// @Description Send event `get_chats` (no payload) to receive chats list (friend + last message).
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ws/get_chats [get]
func GetChatsDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_chats' event."})
}

// MarkMessagesReadDoc godoc
// @Summary Mark messages as read over WebSocket
// @Description Send event `mark_messages_read` with `with_user_id` and optional `up_to_message_id` to set read_at in conversation.
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.MarkMessagesReadRequest true "Mark messages as read payload"
// @Success 200 {object} map[string]string
// @Router /ws/mark_messages_read [post]
func MarkMessagesReadDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'mark_messages_read' event."})
}

// Friend requests
// @Summary Send friend request over WebSocket
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} map[string]string
// @Router /ws/send_friend_request [post]
func SendFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'send_friend_request' event."})
}

// @Summary Accept friend request over WebSocket
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} map[string]string
// @Router /ws/accept_friend_request [post]
func AcceptFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'accept_friend_request' event."})
}

// @Summary Reject friend request over WebSocket
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Friend request payload"
// @Success 200 {object} map[string]string
// @Router /ws/reject_friend_request [post]
func RejectFriendRequestDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'reject_friend_request' event."})
}

// @Summary Remove a friend over WebSocket
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param payload body requests.FriendRequest true "Remove friend payload"
// @Success 200 {object} map[string]string
// @Router /ws/remove_friend [post]
func RemoveFriendDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'remove_friend' event."})
}

// @Summary Get friends list over WebSocket
// @Tags websocket
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ws/get_friends [get]
func GetFriendsDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Docs-only endpoint. Use websocket and send 'get_friends' event."})
}
