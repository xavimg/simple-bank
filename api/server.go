package api

import (
	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repo   db.Store
	router *gin.Engine
}

func NewServer(repo db.Store) *Server {
	server := &Server{
		repo:   repo,
		router: gin.Default(),
	}

	accounts := server.router.Group("/accounts")
	{
		accounts.POST("/", server.createAccount)
		accounts.GET("/:id", server.getAccount)
		accounts.GET("/", server.listAccounts)
	}

	return server
}

func (s *Server) StartRESTServer(add string) error {
	return s.router.Run(add)
}
