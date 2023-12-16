package http

import (
	"metrics/internal/entity"
	"net/http"
	"strconv"
	"strings"
)

var gauge entity.Gauge
var counter entity.Counter

func UpdateGauge(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST-requests are allowed!", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		http.Error(w, "Incorrect request format", http.StatusNotFound)
		return
	}

	name := parts[3]
	if name == "" {
		http.Error(w, "Metric name is required", http.StatusNotFound)
		return
	}

	valueStr := parts[4]
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		http.Error(w, "Incorrect value type", http.StatusBadRequest)
		return
	}

	if gauge.GaugeStorage == nil {
		gauge.GaugeStorage = make(map[string]float64)
	}

	gauge.GaugeStorage[name] = value

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}

func UpdateCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST-requests are allowed!", http.StatusBadRequest)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 {
		http.Error(w, "Metric name is required", http.StatusNotFound)
		return
	}

	name := parts[3]
	if name == "" {
		http.Error(w, "Metric name is required", http.StatusNotFound)
		return
	}
	value, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		http.Error(w, "Incorrect value format", http.StatusBadRequest)
		return
	}

	if counter.CounterStorage == nil {
		counter.CounterStorage = make(map[string]int64)
	}

	counter.CounterStorage[name] += value

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}
