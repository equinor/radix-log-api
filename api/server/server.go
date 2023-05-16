package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Run(ctx context.Context, handler http.Handler, cfg *Config) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler: handler,
	}
	logrus.Infof("starting server with address %s", srv.Addr)
	go startHttp(srv)
	<-ctx.Done()
	return srv.Close()
}

func startHttp(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logrus.Info("server: shutdown complete")
		} else {
			logrus.Errorf("server: %s", err)
		}
	}
}
