package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/mfahrul/elfin-tpl/handler"
	"github.com/mfahrul/elfin-tpl/subscriber"

	elfintpl "github.com/mfahrul/elfin-tpl/proto/elfin-tpl"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.elfin-tpl"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	elfintpl.RegisterElfinTplHandler(service.Server(), new(handler.ElfinTpl))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.elfin-tpl", service.Server(), new(subscriber.ElfinTpl))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
