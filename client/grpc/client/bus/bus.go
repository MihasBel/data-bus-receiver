package bus

import (
	"github.com/MihasBel/data-bus-receiver/client/grpc/gen/v1/bus"
	"github.com/MihasBel/data-bus-receiver/internal/receiver"
)

// Server receiver
type Server struct {
	bus.UnimplementedBusServiceServer
	r receiver.Receiver
}

// New constructor
func New(r receiver.Receiver) *Server {
	return &Server{
		r: r,
	}
}
