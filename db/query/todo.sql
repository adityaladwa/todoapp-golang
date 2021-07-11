-- name: CreateTodo :one
insert into todos (
  title, 
  description
) values (
  $1, $2
) returning *;

-- name: GetTodo :one
select * from todos 
where id=$1 
limit 1;

-- name: ListTodos :many
select * from todos
order by title
limit $1 OFFSET $2;

-- name: DeleteTodo :exec
delete from todos
where id=$1;