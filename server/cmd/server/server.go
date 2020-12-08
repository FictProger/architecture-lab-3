package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/FictProger/architecture-lab-3/server/dormitories"
)

type HttpPortNumber int

// DirectionApiServer configures necessary handlers and starts listening on a configured port.
type DirectionApiServer struct {
	Port HttpPortNumber

	DormitoriesHandler dormitories.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *DirectionApiServer) Start() error {
	if s.DormitoriesHandler == nil {
		return fmt.Errorf("dormitories HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/dormitories", s.DormitoriesHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *DirectionApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
