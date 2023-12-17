package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_ErrServ(t *testing.T) {
	server := New()

	server.Router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "it works")
	})

	httpTest := httptest.NewServer(server.Router)
	defer httpTest.Close()

	res, err := http.Get(httpTest.URL + "/test")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
}
