package api

import (
	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/gin-gonic/gin"
)

// A struct that represents a server instance
type Server struct {
	store  *db.Store
	router *gin.Engine
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

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	apiV1 := router.Group("api/v1")
	apiV1.GET("/todos", server.ListTodos)

	server.router = router
	return server
}

func (server *Server) Start(address string) {
	server.router.Run(address)
}
