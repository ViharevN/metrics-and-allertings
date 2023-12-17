package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetValue(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		r := gin.Default()
		r.GET("/value/:type/:name", GetValue)

		gauge.GaugeStorage["GCCPUFraction"] = 42.5

		req, _ := http.NewRequest(http.MethodGet, "/value/gauge/GCCPUFraction", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "42.5", resp.Body.String())
	})

	t.Run("Invalid Metric", func(t *testing.T) {
		r := gin.Default()
		r.GET("/value/:type/:name", GetValue)

		req, _ := http.NewRequest(http.MethodGet, "/value/gauge/nonexistent", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("Invalid Type", func(t *testing.T) {
		r := gin.Default()
		r.GET("/value/:type/:name", GetValue)

		req, _ := http.NewRequest(http.MethodGet, "/value/invalid/test", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})
}
