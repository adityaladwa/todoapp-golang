package api

import (
	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

// A struct that represents a server instance
type Server struct {
	store  *db.Store
	router *fiber.App
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	app := fiber.New()

	app.Get("v1/api/todos", server.GetTodos)

	server.router = app
	return server
}

func (server *Server) Start(address string) {
	server.router.Listen(address)
}
