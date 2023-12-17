package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"metrics/internal/http"
	"metrics/pkg/httpserver"
)

func Run() {
	//usecase
	//handlers

	//http server
	log.Println("Server started on: localhost:8080")
	server := httpserver.New()
	http.NewRouter(server.Router)
	gin.SetMode(gin.DebugMode)
	server.Start(`localhost:8080`)

}
