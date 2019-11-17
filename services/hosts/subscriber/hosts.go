package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"
    hosts "fastssh/services/hosts/proto/hosts"
)

type HostsHandler struct{}

func (e *HostsHandler) Handle(ctx context.Context, msg *hosts.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *hosts.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}

type hosts struct {

}