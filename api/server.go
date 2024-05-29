package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sawilkhan/simple-bank/db/sqlc"
)

//Server serves HTTP requests for our banking service
type Server struct{
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setd up routing
func NewServer(store *db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/", server.listAccount)

	server.router = router
	return server
}

//Start the HTTP Server on a specific address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}