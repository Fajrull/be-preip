// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"mediversepreip/internal/controllers"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWorldController{},
		},
	})
}
