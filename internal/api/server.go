package api

import "net/http"

type Server struct {
	Endpoint string
	Mux      *http.ServeMux
}

type ServerConfig struct {
	Endpoint string
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		Endpoint: cfg.Endpoint,
	}
}

func (s *Server) RegisterHTTP() {
	s.Mux = http.NewServeMux()

	s.Mux.HandleFunc(`/update/`, s.UpdateHandler)
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(`:8080`, s.Mux)
}
