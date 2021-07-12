package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	db "github.com/adityaladwa/todoapp/db/sqlc"
)

type listTodoRequestQuery struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) ListTodos(c *gin.Context) {
	var req listTodoRequestQuery
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
		return
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

type createTodoRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description,omitempty"`
}

func (s *Server) CreateTodo(c *gin.Context) {
	var req createTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errorResponse(err.Error()))
		return
	}
	var valid bool
	if req.Description == "" {
		valid = false
	} else {
		valid = true
	}
	args := db.CreateTodoParams{
		Title: req.Title,
		Description: sql.NullString{
			String: req.Description,
			Valid:  valid,
		},
	}

	todo, err := s.store.CreateTodo(c.Request.Context(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, mapToResponse(todo))
}
