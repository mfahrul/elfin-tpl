package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	elfintpl "github.com/mfahrul/elfin-tpl/proto/elfin-tpl"
)

type ElfinTpl struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ElfinTpl) Call(ctx context.Context, req *elfintpl.Request, rsp *elfintpl.Response) error {
	log.Info("Received ElfinTpl.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ElfinTpl) Stream(ctx context.Context, req *elfintpl.StreamingRequest, stream elfintpl.ElfinTpl_StreamStream) error {
	log.Infof("Received ElfinTpl.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&elfintpl.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ElfinTpl) PingPong(ctx context.Context, stream elfintpl.ElfinTpl_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&elfintpl.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
