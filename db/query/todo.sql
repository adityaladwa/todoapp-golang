-- name: CreateTodo :one
INSERT INTO todos (
  title, 
  description
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos 
WHERE id=$1 
LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY title
LIMIT $1 OFFSET $2;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id=$1;