-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: GetAllActiveTodos :many
SELECT * FROM todos
WHERE status = 'Active';

-- name: GetAllCompletedTodos :many
SELECT * FROM todos
WHERE status = 'Completed';

-- name: CreateTodo :one
INSERT INTO todos (task, status)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteCompletedTodos :exec
DELETE FROM todos
WHERE status = 'Completed';