package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	db "github.com/adityaladwa/todoapp/db/sqlc"
)

type todoResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type getTodoRequestUri struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) GetTodos(c *gin.Context) {
	var req getTodoRequestUri
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errorResponse(err.Error()))
		return
	}
	args := db.ListTodosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	todos, err := s.store.ListTodos(c.Request.Context(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
	todoResponse := []todoResponse{}
	for i := 0; i < len(todos); i++ {
		todoResponse = append(todoResponse, mapToResponse(todos[i]))
	}
	apiResponse := apiResponse{
		Data:    todoResponse,
		Error:   nil,
		Success: true,
	}
	c.JSON(http.StatusOK, apiResponse)
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
