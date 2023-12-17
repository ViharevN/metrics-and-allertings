package http

import (
	"github.com/gin-gonic/gin"
	"metrics/internal/entity"
	"net/http"
	"strconv"
)

var gauge entity.Gauge
var gaugeStorage = make(map[string]float64)
var counter entity.Counter
var counterStorage = make(map[string]int64)

func UpdateGauge(c *gin.Context) {
	name := c.Param("name")
	valueStr := c.Param("value")

	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric name is required"})
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect value type"})
		return
	}

	gaugeStorage[name] = value

	c.String(http.StatusOK, "OK")
}

func UpdateCounter(c *gin.Context) {
	name := c.Param("name")
	valueStr := c.Param("value")

	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric name is required"})
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect value format"})
		return
	}

	counterStorage[name] += value

	c.String(http.StatusOK, "OK")
}

func GetValue(c *gin.Context) {
	typ := c.Param("type")
	name := c.Param("name")

	var value float64
	var ok bool

	if typ == "gauge" {
		value, ok = gaugeStorage[name]
	} else if typ == "counter" {
		var intValue int64
		intValue, ok = counterStorage[name]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Metric not found"})
			return
		}
		value = float64(intValue)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid metric type"})
		return
	}

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Metric not found"})
		return
	}

	c.String(http.StatusOK, strconv.FormatFloat(value, 'f', -1, 64))
}

func MetricsList(c *gin.Context) {
	metrics := make(map[string]float64)

	for name, value := range gaugeStorage {
		metrics["gauge:"+name] = value
	}

	for name, value := range counterStorage {
		metrics["counter:"+name] = float64(value)
	}

	c.HTML(http.StatusOK, "metrics.tmpl", gin.H{
		"metrics": metrics,
	})
}

func ErrHandler(c *gin.Context) {
	typ := c.Param("type")
	if typ != "counter" && typ != "gauge" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Metric type is required"})
	}
}
