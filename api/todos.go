package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type todoResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

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
