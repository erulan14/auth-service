package model

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type VerifyRequest struct {
	Token string `json:"token" binding:"required"`
}
