package http

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	r.POST("/update/gauge/:name/:value", UpdateGauge)
	r.POST("/update/counter/:name/:value", UpdateCounter)
	r.POST("/update/:type/:name/:value", ErrHandler)
	r.GET("/value/:type/:name", GetValue)
	r.GET("/", MetricsList)
	r.LoadHTMLGlob("pkg/html/template/metrics.tmpl")
}
