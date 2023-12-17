package app

import (
	"log"
	router "metrics/internal/http"
	"metrics/pkg/consts"
	"metrics/pkg/httpserver"
	"net/http"
)

func Run() {
	//usecase
	//handlers
	mux := http.NewServeMux()
	router.NewRouter(mux)
	//http server
	log.Printf("Server started on: %s", consts.Addr)
	server := httpserver.New(mux)

	if err := <-server.ErrServ(); err != nil {
		panic(err)
	}

}
