package app

import (
	"log"
	router "metrics/internal/http"
	"metrics/pkg/httpserver"
	"net/http"
)

func Run() {
	//usecase
	//handlers
	mux := http.NewServeMux()
	router.NewRouter(mux)
	//http server
	log.Println("Server started on: localhost:8080")
	server := httpserver.New(mux)

	if err := <-server.ErrServ(); err != nil {
		panic(err)
	}

}
