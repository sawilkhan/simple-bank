package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/sawilkhan/simple-bank/db/sqlc"
	"github.com/sawilkhan/simple-bank/token"
	"github.com/sawilkhan/simple-bank/util"
)

//Server serves HTTP requests for our banking service
type Server struct{
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setd up routing
func NewServer(config util.Config,store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil{
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}


	server := &Server{
		store: store,
		tokenMaker: tokenMaker,
	}
	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account/", server.listAccount)

	router.POST("/transfer", server.createTransfer)

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	server.router = router
}


//Start the HTTP Server on a specific address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}