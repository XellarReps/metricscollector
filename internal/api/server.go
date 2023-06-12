package api

import (
	"net/http"

	"github.com/XellarReps/metricscollector/internal/storage"
)

type Server struct {
	Address string
	Mux     *http.ServeMux
	Storage storage.Repository
}

type ServerConfig struct {
	Address string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		Address: cfg.Address,
		Storage: storage.Repository(storage.NewMemStorage()),
	}
}

func (s *Server) RegisterHTTP() {
	s.Mux = http.NewServeMux()

	s.Mux.HandleFunc(`/update/`, s.UpdateHandler)
}

func (s *Server) RunServer() error {
	return http.ListenAndServe(s.Address, s.Mux)
}
