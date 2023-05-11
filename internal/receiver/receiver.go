package receiver

import (
	"context"
	"github.com/MihasBel/data-bus/broker/model"
)

type Receiver interface {
	ReceiveMsg(ctx context.Context, message model.Message) error
}
