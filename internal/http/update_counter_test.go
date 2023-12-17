package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateCounter(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/update/counter/test/425", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		UpdateCounter(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		assert.Equal(t, "text/plain", recorder.Header().Get("Content-Type"))
	})

	t.Run("Invalid Method", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/update/counter/test/425", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		UpdateCounter(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("Invalid Metric", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/update/counterss//425", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		UpdateCounter(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("Invalid Value", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/update/counter/test/invalid", nil)
		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		UpdateCounter(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}
