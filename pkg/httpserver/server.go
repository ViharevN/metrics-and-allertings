package httpserver

import (
	"metrics/pkg/consts"
	"net/http"
)

type Server struct {
	server    *http.Server
	errServer chan error
}

// new Server
func New(handler http.Handler) *Server {
	httpserver := &http.Server{
		Handler: handler,
		Addr:    consts.Addr,
	}

	server := &Server{
		server:    httpserver,
		errServer: make(chan error),
	}

	server.start()

	return server
}

// start server
func (s *Server) start() {
	go func() {
		s.errServer <- s.server.ListenAndServe()
		close(s.errServer)
	}()
}

// error started
func (s *Server) ErrServ() <-chan error {
	return s.errServer
}
