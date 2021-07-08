package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/adityaladwa/todoapp/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTodoTx(t *testing.T) {
	store := NewStore(testDB)
	createTodoParam := CreateTodoParams{
		Title: util.RandomTodoTitle(),
		Description: sql.NullString{
			String: util.RandomTodoDescription(),
			Valid:  true,
		},
	}

	// run n concurrent transactions
	n := 100

	errs := make(chan error)
	results := make(chan Todo)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateTodo(context.Background(), createTodoParam)
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
		require.Equal(t, createTodoParam.Title, result.Title)
		require.Equal(t, result.Description, result.Description)

		_, err = store.GetTodo(context.Background(), result.ID)
		require.NoError(t, err)
	}
}
