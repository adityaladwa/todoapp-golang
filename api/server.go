package api

import (
	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// A struct that represents a server instance
type Server struct {
	store  *db.Store
	router *fiber.App
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	v1 := app.Group("v1/api")
	
	v1.Get("/todos", server.GetTodos)

	server.router = app
	return server
}

func (server *Server) Start(address string) {
	server.router.Listen(address)
}
