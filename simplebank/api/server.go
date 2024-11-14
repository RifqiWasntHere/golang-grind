package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Newserver creates a new  HTTP server and setups routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccounts)
	server.router = router
	return server
}

// starts http server on given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// basically just forward the error in map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
