package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateGauge(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		r := gin.Default()
		r.POST("/update/gauge/:name/:value", UpdateGauge)

		req, _ := http.NewRequest(http.MethodPost, "/update/gauge/test/42.5", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, "text/plain; charset=utf-8", resp.Header().Get("Content-Type"))
	})

	t.Run("Invalid Method", func(t *testing.T) {
		r := gin.Default()
		r.POST("/update/gauge/:name/:value", UpdateGauge)

		req, _ := http.NewRequest(http.MethodGet, "/update/gauge/test/42.5", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("Invalid Metric", func(t *testing.T) {
		r := gin.Default()
		r.POST("/update/gauge/:name/:value", UpdateGauge)

		req, _ := http.NewRequest(http.MethodPost, "/update/gauge//42.5", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusNotFound, resp.Code)
	})

	t.Run("Invalid Value", func(t *testing.T) {
		r := gin.Default()
		r.POST("/update/gauge/:name/:value", UpdateGauge)

		req, _ := http.NewRequest(http.MethodPost, "/update/gauge/test/invalid", nil)
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}
