package client

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/MihasBel/data-bus-receiver/client/grpc/client/bus"
	v1bus "github.com/MihasBel/data-bus-receiver/client/grpc/gen/v1/bus"
	"github.com/MihasBel/data-bus-receiver/internal/receiver"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// MSGs
const (
	TCP = "tcp"

	MsgStartListening = "start listening grpc at %v"
	MsgStopListening  = "stop listening grpc at %v"
	MsgServerFailed   = "client failed"

	MsgErrFailedListen = "failed to listen GRPC client: %w"

	KeyLoggerDirection  = "direction"
	KeyLoggerGRPCStatus = "grpc_status"
	KeyLoggerDuration   = "duration"
	KeyLoggerAnswer     = "answer"
	KeyLoggerError      = "error"

	ValLoggerDirection = "in"
)

// Server represents grpc client
type Server struct {
	cfg Config
	srv *grpc.Server
	r   receiver.Receiver
	l   zerolog.Logger
}

// New Server constructor
func New(
	cfg Config,
	r receiver.Receiver,
	l zerolog.Logger,
) *Server {

	return &Server{
		cfg: cfg,
		r:   r,
		l:   l,
	}
}

// Start client
func (s *Server) Start(_ context.Context) error {
	lis, err := net.Listen(TCP, s.cfg.Host)
	if err != nil {
		return errors.Wrap(err, MsgErrFailedListen)
	}

	s.srv = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			s.ZerologInterceptor,
			recovery.UnaryServerInterceptor(),
		)),
	)
	v1bus.RegisterBusServiceServer(s.srv, bus.New(s.r))

	reflection.Register(s.srv)

	s.l.Info().Msgf(MsgStartListening, s.cfg.Host)
	errCh := make(chan error)
	go func() {
		if err := s.srv.Serve(lis); err != nil {
			s.l.Error().Err(err).Msg(MsgServerFailed)
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Duration(s.cfg.StartTimeout) * time.Second):
		return nil
	}
}

// Stop client
func (s *Server) Stop(_ context.Context) error {
	s.l.Info().Msgf(MsgStopListening, s.cfg.Host)
	stopCh := make(chan struct{})
	go func() {
		s.srv.GracefulStop()
		stopCh <- struct{}{}
	}()
	select {
	case <-time.After(time.Duration(s.cfg.StopTimeout) * time.Second):
		return nil
	case <-stopCh:
		return nil
	}
}

// ZerologInterceptor logger
func (s *Server) ZerologInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	start := time.Now()
	resp, err := handler(ctx, req)
	body, _ := json.Marshal(resp)
	st, _ := status.FromError(err)

	lgr := s.l.Info().
		Str(KeyLoggerDirection, ValLoggerDirection).
		Str(KeyLoggerGRPCStatus, st.Code().String()).
		Dur(KeyLoggerDuration, time.Since(start)).
		Str(KeyLoggerAnswer, string(body))

	if err != nil {
		lgr.Str(KeyLoggerError, err.Error())
	}

	lgr.Msg(info.FullMethod)

	return resp, err
}
