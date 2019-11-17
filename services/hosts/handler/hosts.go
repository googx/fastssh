package handler

import (
	"context"
	hosts "fastssh/services/hosts/proto/hosts"
	
	"github.com/micro/go-micro/util/log"
	
	
)

type Hosts struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Hosts) Call(ctx context.Context, req *hosts.Request, rsp *hosts.Response) error {
	
	log.Log("Received Hosts.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Hosts) Stream(ctx context.Context, req *hosts.StreamingRequest, stream hosts.Hosts_StreamStream) error {
	log.Logf("Received Hosts.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&hosts.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Hosts) PingPong(ctx context.Context, stream hosts.Hosts_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&hosts.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
