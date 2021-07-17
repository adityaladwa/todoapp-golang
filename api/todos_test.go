package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/adityaladwa/todoapp/db/mock"
	"github.com/stretchr/testify/require"

	db "github.com/adityaladwa/todoapp/db/sqlc"
	"github.com/adityaladwa/todoapp/util"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestGetTodoApi(t *testing.T) {
	todo := randomTodo()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mockdb.NewMockQuerier(ctrl)

	// build stubs
	mockQuerier.EXPECT().
		GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
		Times(1).
		Return(todo, nil)

	// start a test http server
	server := NewServer(mockQuerier)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/api/v1/todos/%s", todo.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, todo)
}

func randomTodo() db.Todo {
	now, err := time.Parse(time.RFC3339Nano, "2021-07-13T10:23:41.40568+05:30")
	if err != nil {
		log.Fatal(err)
	}
	return db.Todo{
		ID:    uuid.New(),
		Title: util.RandomTodoTitle(),
		Description: sql.NullString{
			String: util.RandomTodoDescription(),
			Valid:  true,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, todo db.Todo) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	expected := mapTodoResponse(todo)

	var gotTodo todoResponse
	err = json.Unmarshal(data, &gotTodo)

	require.NoError(t, err)
	require.Equal(t, expected, gotTodo)
}
