package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (s *Store) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	var todo Todo
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		todo, err = q.CreateTodo(ctx, arg)
		if err != nil {
			return err
		}

		return nil
	})
	return todo, err
}

func (s *Store) GetTodo(ctx context.Context, id uuid.UUID) (Todo, error) {
	var todo Todo
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		todo, err = q.GetTodo(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	return todo, err
}
