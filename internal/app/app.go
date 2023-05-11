package app

import (
	"context"
	"github.com/MihasBel/data-bus-receiver/pkg/lifecycle"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var (
	ErrStartTimeout    = errors.New("start timeout")
	ErrShutdownTimeout = errors.New("shutdown timeout")
)

type (
	App struct {
		log  *zerolog.Logger
		cmps []cmp
		cfg  Config
	}
	cmp struct {
		Service lifecycle.Lifecycle
		Name    string
	}
)

func New(cfg Config, l zerolog.Logger) *App {
	l = l.With().Str("cmp", "app").Logger()

	return &App{
		log: &l,
		cfg: cfg,
	}
}

func (a *App) Start(ctx context.Context) error {
	a.log.Info().Msg("starting app")

	a.cmps = append(
		a.cmps,
		cmp{grpcCli, "grpcClient"},
		cmp{b, "broker"},
		cmp{server, "server"},
	)

	okCh, errCh := make(chan struct{}), make(chan error)

	go func() {
		for _, c := range a.cmps {
			a.log.Info().Msgf("%v is starting", c.Name)

			if err := c.Service.Start(ctx); err != nil {
				a.log.Error().Err(err).Msgf(FmtCannotStart, c.Name)
				errCh <- errors.Wrapf(err, FmtCannotStart, c.Name)

				return
			}

			a.log.Info().Msgf("%v started", c.Name)
		}
		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ErrStartTimeout
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info().Msg("Application started!")
		return nil
	}
}
