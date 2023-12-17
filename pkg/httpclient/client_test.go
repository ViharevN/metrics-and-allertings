package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUpdateMetric(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		expectedURL := "/update/gauge/TestMetric/123.45"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != http.MethodPost {
				t.Errorf("Expected method 'POST', got %s", r.Method)
			}

			if r.URL.EscapedPath() != expectedURL {
				t.Errorf("Expected URL path '%s', got '%s'", expectedURL, r.URL.EscapedPath())
			}

			if r.Header.Get("Content-Type") != "text/plain" {
				t.Errorf("Expected header 'Content-Type: text/plain'")
			}

			w.WriteHeader(http.StatusOK)
		}))

		defer server.Close()

		agent := NewAgent(server.URL, 2*time.Second, 10*time.Second)
		agent.sendMetric("gauge", "TestMetric", 123.45)
	})
}
