-- name: CreateUser :exec
INSERT INTO users17 (username, email,password) VALUES ($1, $2, $3);

-- name: GetUser :one
SELECT * FROM users17 WHERE email=$1;

-- name: CreateTask :one
INSERT INTO tasks17 (user_id,title,created_at) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks17
WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: ListTasks :many
SELECT * FROM tasks17
WHERE user_id = $1 ORDER BY id ASC;

-- name: UpdateTask :one
UPDATE tasks17
  set title = $3,
  updated_at = $4
WHERE id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks17
WHERE id = $1 AND user_id = $2;