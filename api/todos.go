package api

import (
	"net/http"
	"strconv"
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
	pageSize, err := strconv.ParseInt(c.Query("page_size"), 10, 32)
	if err != nil {
		c.SendStatus(http.StatusUnprocessableEntity)
	}
	pageNo, err := strconv.ParseInt(c.Query("page_no"), 10, 32)
	if err != nil {
		c.SendStatus(http.StatusUnprocessableEntity)
	}
	args := db.ListTodosParams{
		Limit:  int32(pageSize),
		Offset: int32(pageNo),
	}
	todos, err := s.store.ListTodos(c.UserContext(), args)
	if err != nil {
		c.SendString("error")
		return c.SendStatus(http.StatusInternalServerError)
	}

	var todoResponse []todoResponse
	for i := 0; i < len(todos); i++ {
		todoResponse = append(todoResponse, mapToResponse(todos[i]))
	}
	return c.JSON(todoResponse)
}

func mapToResponse(t db.Todo) todoResponse {
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
