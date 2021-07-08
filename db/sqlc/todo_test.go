package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/adityaladwa/todoapp/util"
	"github.com/stretchr/testify/require"
)

func CreateTodo(t *testing.T) Todo {
	arg := CreateTodoParams{
		Title: util.RandomTodoTitle(),
		Description: sql.NullString{
			String: util.RandomTodoDescription(),
			Valid:  true,
		},
	}

	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, todo.Description.String, todo.Description.String)

	require.NotEmpty(t, todo.ID)
	require.NotEmpty(t, todo.CreatedAt)
	require.NotEmpty(t, todo.UpdatedAt)

	return todo
}
func TestCreateTodo(t *testing.T) {
	CreateTodo(t)
}

func TestGetTodo(t *testing.T) {
	todo1 := CreateTodo(t)
	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.NoError(t, err)

	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Title, todo2.Title)
	require.Equal(t, todo1.Description, todo2.Description)
	require.WithinDuration(t, todo1.CreatedAt, todo2.CreatedAt, time.Second)
	require.WithinDuration(t, todo1.UpdatedAt, todo2.UpdatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	todo1 := CreateTodo(t)
	err := testQueries.DeleteTodo(context.Background(), todo1.ID)
	require.NoError(t, err)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, todo2)
}
