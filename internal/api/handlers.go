package api

import (
	"bytes"
	"html/template"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	gaugeType   = "gauge"
	counterType = "counter"
)

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Storage.UpdateStorage(chi.URLParam(r, "type"), chi.URLParam(r, "name"), chi.URLParam(r, "value"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getMetricHandler(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "type")
	if metricType != counterType && metricType != gaugeType {
		http.Error(w, "unknown metric type", http.StatusBadRequest)
		return
	}

	val, err := s.Storage.GetMetricFromStorage(metricType, chi.URLParam(r, "name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val))
}

func (s *Server) getAllMetricsHandler(w http.ResponseWriter, r *http.Request) {
	gauge, counter := s.Storage.GetAllMetricsFromStorage()

	tpl, err := template.ParseFiles("internal/api/templates/answer.html")
	if err != nil {
		http.Error(w, "cannot open html template", http.StatusInternalServerError)
		return
	}

	data := make(map[string]any)
	data["gauge"] = gauge
	data["counter"] = counter

	var buf bytes.Buffer

	if err = tpl.Execute(&buf, data); err != nil {
		http.Error(w, "cannot execute template", http.StatusInternalServerError)
		return
	}

	page, err := io.ReadAll(&buf)
	if err != nil {
		http.Error(w, "cannot read buffer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "Content-Type: text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(page)
}
