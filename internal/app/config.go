package app

import (
	"github.com/MihasBel/data-bus-receiver/client/grpc/client"
	"github.com/MihasBel/data-bus-receiver/delivery/kafka"
	"time"
)

type Config struct {
	LogLevel    string   `env:"LOG_LEVEL" envDefault:"info"`
	MsgTypes    []string `env:"MSG_TYPES" required:"true"`
	KafkaConfig kafka.Config
	GRPCConfig  client.Config

	StartTimeout time.Duration `env:"START_TIMEOUT"`
	StopTimeout  time.Duration `env:"STOP_TIMEOUT"`
}
