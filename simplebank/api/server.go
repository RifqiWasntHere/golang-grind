package api

import (
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Newserver creates a new  HTTP server and setups routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// binding.Validator.Engine()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validateCurrency)
	}

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccounts)

	router.POST("/transfer", server.createTransfer)
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
