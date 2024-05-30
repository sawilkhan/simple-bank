package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/sawilkhan/simple-bank/db/sqlc"
)

//Server serves HTTP requests for our banking service
type Server struct{
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setd up routing
func NewServer(store db.Store) *Server{
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/", server.listAccount)

	router.POST("/transfer", server.createTransfer)

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