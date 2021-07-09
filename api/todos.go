package api

import (
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

func mapFromTodos(t []db.Todo) []todoResponse {
	r := make([]todoResponse, len(t))
	for i := 0; i < len(t); i++ {
		r[i].ID = t[i].ID
		r[i].Title = t[i].Title
		if t[i].Description.Valid {
			r[i].Description = t[i].Description.String
		} else {
			r[i].Description = ""
		}
		r[i].CreatedAt = t[i].CreatedAt
		r[i].UpdatedAt = t[i].UpdatedAt
	}
	return r
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
	return c.JSON(mapFromTodos(todos))
}
