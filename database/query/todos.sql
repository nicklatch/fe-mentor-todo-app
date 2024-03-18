-- name: GetAllTodos :many
SELECT *
FROM todos
ORDER BY id;

-- name: GetAllActiveTodos :many
SELECT *
FROM todos
WHERE completed = false
ORDER BY id;

-- name: GetAllCompletedTodos :many
SELECT *
FROM todos
WHERE completed = true
ORDER BY id;

-- name: CreateTodo :one
INSERT INTO todos (task, completed)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteOneTodo :exec
DELETE FROM todos
WHERE id = $1;

-- name: DeleteCompletedTodos :exec
DELETE
FROM todos
WHERE completed = true;

-- name: ToggleCompleted :exec
UPDATE todos
SET completed = NOT completed, updated_at = now()
WHERE id = $1;
