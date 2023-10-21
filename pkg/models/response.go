package models

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewErrorResponse(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func NewInternalServerError() *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
}

type SuccessResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewSuccessResponse(statusCode int, message string) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
