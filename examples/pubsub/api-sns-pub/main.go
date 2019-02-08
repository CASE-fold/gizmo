package main

import (
	"github.com/CASE-fold/gizmo/v2/config"
	"github.com/CASE-fold/gizmo/v2/examples/pubsub/api-sns-pub/service"
	"github.com/CASE-fold/gizmo/v2/pubsub/aws"
	"github.com/CASE-fold/gizmo/v2/server"
)

func main() {
	var cfg struct {
		Server *server.Config
		SNS    aws.SNSConfig
	}
	config.LoadJSONFile("./config.json", &cfg)

	server.Init("nyt-json-pub-proxy", cfg.Server)

	err := server.Register(service.NewJSONPubService(cfg.SNS))
	if err != nil {
		server.Log.Fatal("unable to register service: ", err)
	}

	err = server.Run()
	if err != nil {
		server.Log.Fatal("server encountered a fatal error: ", err)
	}
}
