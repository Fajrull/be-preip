package main

import (
	"github.com/sev-2/raiden"
	"mediversepreip/internal/bootstrap"
)

func main() {
	// load configuration
	config, err := raiden.LoadConfig(nil)
	if err != nil {
		raiden.Error("load configuration",err.Error())
	}

	// Setup server
	server := raiden.NewServer(config)

	// register route
	bootstrap.RegisterRoute(server)
	bootstrap.RegisterJobs(server)

	// run server
	server.Run()
}
