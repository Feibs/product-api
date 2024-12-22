package dto

type MessageResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
