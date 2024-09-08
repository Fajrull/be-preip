package controllers

import (
	"github.com/sev-2/raiden"
)

type HelloWorldRequest struct { // Add your request data 
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

type HelloWorldController struct {
	raiden.ControllerBase
	Http	string `path:"/hello" type:"custom"`
	Payload	*HelloWorldRequest
	Result	HelloWorldResponse
}

func (c *HelloWorldController) Get(ctx raiden.Context) error {
	c.Result.Message = "Welcome to raiden"
	return ctx.SendJson(c.Result)
}
