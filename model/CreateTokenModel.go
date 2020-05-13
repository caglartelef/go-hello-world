package model

type CreateTokenRequest struct {
	Username string `form:"username" validate:"required_all=Password" qs:"username"`
	Password string `form:"password" validate:"required_all=Username" qs:"password"`
}
