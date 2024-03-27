package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, handler http.Handler, cfg *Config) error {
	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:     handler,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	log.Ctx(ctx).Info().Msgf("API is serving on address %s", srv.Addr)
	go startHttp(ctx, srv)
	<-ctx.Done()
	return srv.Close()
}

func startHttp(ctx context.Context, s *http.Server) {
	logger := log.Ctx(ctx)
	if err := s.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Info().Msg("server closed")
		} else {
			logger.Fatal().Err(err).Msg("server error")
		}
	}
}
