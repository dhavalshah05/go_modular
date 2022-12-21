package models

type ApiResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ErrorApiResponse(message string) ApiResponse {
	return ApiResponse{
		Message: message,
		Data:    nil,
	}
}
