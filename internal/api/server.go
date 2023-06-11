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
	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, UpdateHandler)

	return &Server{
		Endpoint: cfg.Endpoint,
		Mux:      mux,
	}
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(`:8080`, s.Mux)
}
