package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gofiber/fiber/v2"
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

	var todoResponse []todoResponse
	for i := 0; i < len(todos); i++ {
		todoResponse = append(todoResponse, mapTodoToResponse(todos[i]))
	}
	return c.JSON(todoResponse)
}

func mapTodoToResponse(t db.Todo) todoResponse {
	var description string
	if t.Description.Valid {
		description = t.Description.String
	} else {
		description = ""
	}
	return todoResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}