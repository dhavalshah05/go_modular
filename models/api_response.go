package models

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
