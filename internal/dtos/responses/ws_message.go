package responses

type ResponseMessage struct {
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
