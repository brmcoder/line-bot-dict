package controller

import (
	"github.com/brmcoder/line-bot-dict/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	router *gin.Engine
}

func NewServer(config util.Config) *Server {
	server := &Server{config: config}
	router := gin.Default()

	server.NewWebhookController(router)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run()
	// return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
