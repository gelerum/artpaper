package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) (err error) {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	err = s.httpServer.ListenAndServe()
	if err != nil {
		log.Printf("Error occurred while starting server: %s", err)
	}
	return
}

func (s *Server) Shutdown(ctx context.Context) (err error) {
	err = s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Error occurred while shutting down server: %s", err)
	}
	return
}
