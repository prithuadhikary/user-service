package model

type SignupRequest struct {
	Username             string `json:"username" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}
