package model

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseToken struct {
	Token string `json:"token"`
}
