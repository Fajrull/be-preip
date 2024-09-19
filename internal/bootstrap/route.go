// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"github.com/valyala/fasthttp"
	"mediversepreip/internal/controllers"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/auth/v1/token?grant_type=password",
			Methods:    []string{fasthttp.MethodGet, fasthttp.MethodPost},
			Controller: &controllers.AuthController{},
		},
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWorldController{},
		},
	})
}
