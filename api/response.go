package api

import (
	"time"

	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/google/uuid"
)

type todoResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func mapTodoResponse(t db.Todo) todoResponse {
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

// Generic API respose
type apiResponse struct {
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Success bool        `json:"success"`
}

// Generic API error response
type apiErrorResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// A function that returns a default error
func errorResponse(error string) apiResponse {
	errorResponse := apiErrorResponse{
		Title:   "Oops, something went wrong",
		Message: error,
		Code:    "error.server.db.",
	}
	return apiResponse{
		Data:    nil,
		Error:   errorResponse,
		Success: false,
	}
}
