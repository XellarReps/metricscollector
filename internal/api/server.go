package api

import (
	"net/http"

	"github.com/XellarReps/metricscollector/internal/storage"
)

type Server struct {
	Endpoint string
	Mux      *http.ServeMux
	Storage  *storage.MemStorage
}

type ServerConfig struct {
	Endpoint string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		Endpoint: cfg.Endpoint,
		Storage:  storage.NewMemStorage(),
	}
}

func (s *Server) RegisterHTTP() {
	s.Mux = http.NewServeMux()

	s.Mux.HandleFunc(`/update/`, s.UpdateHandler)
}

func (s *Server) RunServer() error {
	return http.ListenAndServe(s.Endpoint, s.Mux)
}
