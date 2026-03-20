package responses

import "github.com/google/uuid"

type GroupResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}