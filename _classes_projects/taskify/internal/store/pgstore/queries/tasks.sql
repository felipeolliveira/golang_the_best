-- name: CreateTask :one
INSERT INTO tasks (
  title, description, priority
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTaskById :one
SELECT *
FROM tasks WHERE id = $1;

-- name: ListTasks :many
SELECT *
FROM tasks
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, description = $3, priority = $4, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = $1;
