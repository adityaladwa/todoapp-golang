package api

import (
	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gin-gonic/gin"
)

// A struct that represents a server instance
type Server struct {
	querier db.Querier
	router  *gin.Engine
}

func NewServer(querier db.Querier) *Server {
	server := &Server{querier: querier}
	router := gin.Default()

	apiV1 := router.Group("api/v1")
	apiV1.GET("/todos", server.ListTodos)
	apiV1.POST("/todos", server.CreateTodo)
	apiV1.GET("/todos/:id", server.GetTodo)

	server.router = router
	return server
}

func (server *Server) Start(address string) {
	server.router.Run(address)
}
