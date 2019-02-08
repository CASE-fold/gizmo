package main

import (
	"github.com/case-fold/gizmo/examples/servers/mixed/service"

	"github.com/case-fold/gizmo/config"
	"github.com/case-fold/gizmo/server"
)

func main() {
	// showing 1 way of managing gizmo/config: importing from a local file
	var cfg *service.Config
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("nyt-simple-proxy", cfg.Server)

	err := server.Register(service.NewMixedService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("server encountered a fatal error: ", err)
	}
}
