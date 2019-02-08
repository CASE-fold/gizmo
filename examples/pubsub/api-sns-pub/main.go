package main

import (
	"github.com/case-fold/gizmo/config"
	"github.com/case-fold/gizmo/examples/pubsub/api-sns-pub/service"
	"github.com/case-fold/gizmo/pubsub/aws"
	"github.com/case-fold/gizmo/server"
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
