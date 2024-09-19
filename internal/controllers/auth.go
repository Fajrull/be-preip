package controllers

import (
	"github.com/sev-2/raiden"
)

type AuthRequest struct { // Add your request data
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AuthController struct {
	raiden.ControllerBase
	Http    string `path:"/auth/v1/token?grant_type=password" type:"custom"`
	Payload *AuthRequest
	Result  AuthResponse
}

func (c *AuthController) Get(ctx raiden.Context) error {
	c.Result.Message = "(auth) welcome to raiden"
	return ctx.SendJson(c.Result)
}

func (c *AuthController) Post(ctx raiden.Context) error {
	c.Result.Message = "Login successful"
	return ctx.SendJson(c.Result)
}
