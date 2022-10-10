package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	HttpServer *http.Server
}

func (s *Server) Run(post string, handler http.Handler) error {
	s.HttpServer = &http.Server{
		Addr:           ":" + post,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
