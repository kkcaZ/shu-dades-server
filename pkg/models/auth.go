package models

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	StatusCode int        `json:"statusCode"`
	UserClaim  *UserClaim `json:"userClaim"`
}

type UserClaim struct {
	Token string `json:"token"`
	Role  Role   `json:"role"`
}
