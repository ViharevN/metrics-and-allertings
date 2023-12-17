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

			metrics = append(metrics, entity.Metric{"gauge", "Alloc", float64(memStats.Alloc)})
			metrics = append(metrics, entity.Metric{"gauge", "BuckHashSys", float64(memStats.BuckHashSys)})
			metrics = append(metrics, entity.Metric{"gauge", "Frees", float64(memStats.Frees)})
			metrics = append(metrics, entity.Metric{"gauge", "GCCPUFraction", float64(memStats.GCCPUFraction)})
			metrics = append(metrics, entity.Metric{"gauge", "GCSys", float64(memStats.GCSys)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapAlloc", float64(memStats.HeapAlloc)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapIdle", float64(memStats.HeapIdle)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapInuse", float64(memStats.HeapInuse)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapObjects", float64(memStats.HeapObjects)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapReleased", float64(memStats.HeapReleased)})
			metrics = append(metrics, entity.Metric{"gauge", "HeapSys", float64(memStats.HeapSys)})
			metrics = append(metrics, entity.Metric{"gauge", "LastGC", float64(memStats.LastGC)})
			metrics = append(metrics, entity.Metric{"gauge", "Lookups", float64(memStats.Lookups)})
			metrics = append(metrics, entity.Metric{"gauge", "MCacheInuse", float64(memStats.MCacheInuse)})
			metrics = append(metrics, entity.Metric{"gauge", "MCacheSys", float64(memStats.MCacheSys)})
			metrics = append(metrics, entity.Metric{"gauge", "MSpanInuse", float64(memStats.MSpanInuse)})
			metrics = append(metrics, entity.Metric{"gauge", "MSpanSys", float64(memStats.MSpanSys)})
			metrics = append(metrics, entity.Metric{"gauge", "Mallocs", float64(memStats.Mallocs)})
			metrics = append(metrics, entity.Metric{"gauge", "NextGC", float64(memStats.NextGC)})
			metrics = append(metrics, entity.Metric{"gauge", "NumForcedGC", float64(memStats.NumForcedGC)})
			metrics = append(metrics, entity.Metric{"gauge", "NumGC", float64(memStats.NumGC)})
			metrics = append(metrics, entity.Metric{"gauge", "OtherSys", float64(memStats.OtherSys)})
			metrics = append(metrics, entity.Metric{"gauge", "PauseTotalNs", float64(memStats.PauseTotalNs)})
			metrics = append(metrics, entity.Metric{"gauge", "StackInuse", float64(memStats.StackInuse)})
			metrics = append(metrics, entity.Metric{"gauge", "StackSys", float64(memStats.StackSys)})
			metrics = append(metrics, entity.Metric{"gauge", "Sys", float64(memStats.Sys)})
			metrics = append(metrics, entity.Metric{"gauge", "TotalAlloc", float64(memStats.TotalAlloc)})

			metrics = append(metrics, entity.Metric{"gauge", "RandomValue", rand.Float64()})
			metrics = append(metrics, entity.Metric{"counter", "PollCount", 1})

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

	if resp.StatusCode != http.StatusOK {
		fmt.Println("server returned non-OK status:", resp.Status)
	}
}
