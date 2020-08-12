package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	elfintpl "github.com/mfahrul/elfin-tpl/proto/elfin-tpl"
)

type ElfinTpl struct{}

func (e *ElfinTpl) Handle(ctx context.Context, msg *elfintpl.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *elfintpl.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
