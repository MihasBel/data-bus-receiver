package kafka

import (
	"context"
	"github.com/MihasBel/data-bus/broker/model"
	"github.com/pkg/errors"
)

func (b *Broker) ReceiveMsg(_ context.Context, message model.Message) error {
	if !contains(b.msgTypes, message.MsgType) {
		return errors.Errorf(MsgErrUnknownType, message.MsgType)
	}
	if len(message.Data) == 0 {
		return errors.Errorf(MsgErrEmptyMsg, message.MsgType)
	}
	return b.produce(message.MsgType, message.Data)
}

func contains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
			return true
		}
	}
	return false
}
