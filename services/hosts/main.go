package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"

	hosts "hosts/proto/hosts"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("top.goox.srv.hosts"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()

	// Register Handler
	hosts.RegisterHostsHandler(service.Server(), new(handler.Hosts))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("top.goox.srv.hosts", service.Server(), new(subscriber.Hosts))

	// Register Function as Subscriber
	micro.RegisterSubscriber("top.goox.srv.hosts", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
