package httpclient

import (
	"fmt"
	"math/rand"
	"metrics/internal/entity"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type Agent struct {
	serverURL      string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func NewAgent(serverURL string, pollInterval, reportInterval time.Duration) *Agent {
	return &Agent{
		serverURL:      serverURL,
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
	}
}

func (a *Agent) Run() {
	pollTicker := time.NewTicker(a.pollInterval)
	reportTicker := time.NewTicker(a.reportInterval)
	defer pollTicker.Stop()
	defer reportTicker.Stop()

	var metrics []entity.Metric

	for {
		select {
		case <-pollTicker.C:
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)

			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "Alloc", Value: float64(memStats.Alloc)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "BuckHashSys", Value: float64(memStats.BuckHashSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "Frees", Value: float64(memStats.Frees)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "GCCPUFraction", Value: float64(memStats.GCCPUFraction)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "GCSys", Value: float64(memStats.GCSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapAlloc", Value: float64(memStats.HeapAlloc)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapIdle", Value: float64(memStats.HeapIdle)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapInuse", Value: float64(memStats.HeapInuse)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapObjects", Value: float64(memStats.HeapObjects)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapReleased", Value: float64(memStats.HeapReleased)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "HeapSys", Value: float64(memStats.HeapSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "LastGC", Value: float64(memStats.LastGC)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "Lookups", Value: float64(memStats.Lookups)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "MCacheInuse", Value: float64(memStats.MCacheInuse)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "MCacheSys", Value: float64(memStats.MCacheSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "MSpanInuse", Value: float64(memStats.MSpanInuse)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "MSpanSys", Value: float64(memStats.MSpanSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "Mallocs", Value: float64(memStats.Mallocs)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "NextGC", Value: float64(memStats.NextGC)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "NumForcedGC", Value: float64(memStats.NumForcedGC)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "NumGC", Value: float64(memStats.NumGC)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "OtherSys", Value: float64(memStats.OtherSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "PauseTotalNs", Value: float64(memStats.PauseTotalNs)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "StackInuse", Value: float64(memStats.StackInuse)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "StackSys", Value: float64(memStats.StackSys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "Sys", Value: float64(memStats.Sys)})
			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "TotalAlloc", Value: float64(memStats.TotalAlloc)})

			metrics = append(metrics, entity.Metric{Type: "gauge", Name: "RandomValue", Value: rand.Float64()})
			metrics = append(metrics, entity.Metric{Type: "counter", Name: "PollCount", Value: 1})

		case <-reportTicker.C:
			for _, metric := range metrics {
				a.sendMetric(metric.Type, metric.Name, metric.Value)
			}
			metrics = nil
		}
	}
}

func (a *Agent) sendMetric(typ, name string, value float64) {
	url := fmt.Sprintf("%s/update/%s/%s/%s", a.serverURL, typ, name, strconv.FormatFloat(value, 'f', -1, 64))
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		fmt.Println("failed to send metric:", err)
		return
	}
	defer resp.Body.Close()
}
