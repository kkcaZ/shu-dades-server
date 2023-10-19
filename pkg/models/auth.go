package models

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	StatusCode int    `json:"statusCode"`
	Token      string `json:"token"`
}
