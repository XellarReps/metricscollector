package api

import (
	"net/http"
	"strings"
)

const (
	metricType = iota + 2
	metricName
	metricValue
)

func (s *Server) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	elems := strings.Split(r.RequestURI, "/")
	if len(elems) != 5 {
		http.Error(w, "some of the request elements are missing", http.StatusNotFound)
		return
	}

	err := s.Storage.UpdateStorage(elems[metricType], elems[metricName], elems[metricValue])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
