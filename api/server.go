package api

import (
	db "github.com/adamwgriffin/go-microservice/db/sqlc"
	"github.com/adamwgriffin/go-microservice/lib"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our service.
type Server struct {
	config lib.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config lib.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/ping", server.ping)
	router.GET("/contact/:id", server.getContact)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
