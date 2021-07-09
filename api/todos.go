package api

import (
	"net/http"

	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetTodos(c *fiber.Ctx) error {
	args := db.ListTodosParams{
		Limit:  10,
		Offset: 0,
	}

	todos, err := s.store.ListTodos(c.UserContext(), args)
	if err != nil {
		c.SendString("error")
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.JSON(todos)
}
