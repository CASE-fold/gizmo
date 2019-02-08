package main

import (
	"github.com/CASE-fold/gizmo/v2/examples/servers/json/service"

	"github.com/CASE-fold/gizmo/v2/config"
	"github.com/CASE-fold/gizmo/v2/server"
)

func main() {
	// showing 1 way of managing gizmo/config: importing from a local file
	var cfg *service.Config
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("nyt-most-popular-json-proxy", cfg.Server)

	err := server.Register(service.NewJSONService(cfg))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("server encountered a fatal error: ", err)
	}
}
