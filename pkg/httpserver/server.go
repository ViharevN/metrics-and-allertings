package httpserver

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
}

// new Server
func New() *Server {
	r := gin.Default()

	server := &Server{
		Router: r,
	}

	return server
}

// start server
func (s *Server) Start(addr string) {
	s.Router.Run(addr)
}
