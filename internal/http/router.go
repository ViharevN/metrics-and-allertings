package http

import (
	"net/http"
)

func NewRouter(mux *http.ServeMux) {
	mux.HandleFunc("/update/gauge/", UpdateGauge)
	mux.HandleFunc("/update/counter/", UpdateCounter)
	mux.HandleFunc("/update/", ErrValidTypeMetric)
}
