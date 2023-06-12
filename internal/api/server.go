package api

import (
	"net/http"

	"github.com/XellarReps/metricscollector/internal/storage"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Address string
	Mux     *chi.Mux
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
	s.Mux = chi.NewRouter()

	s.Mux.Route("/", func(r chi.Router) {
		r.Get("/", s.getAllMetricsHandler)
		r.Get("/value/{type}/{name}", s.getMetricHandler)
		r.Post("/update/{type}/{name}/{value}", s.updateHandler)
	})
}

func (s *Server) RunServer() error {
	return http.ListenAndServe(s.Address, s.Mux)
}
