// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package storage

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks17 (user_id,title,created_at) VALUES ($1, $2, $3)
RETURNING id, user_id, title, created_at, updated_at
`

type CreateTaskParams struct {
	UserID    sql.NullInt32
	Title     sql.NullString
	CreatedAt sql.NullString
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Tasks17, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.UserID, arg.Title, arg.CreatedAt)
	var i Tasks17
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO users17 (username, email,password) VALUES ($1, $2, $3)
`

type CreateUserParams struct {
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks17
WHERE id = $1 AND user_id = $2
`

type DeleteTaskParams struct {
	ID     int32
	UserID sql.NullInt32
}

func (q *Queries) DeleteTask(ctx context.Context, arg DeleteTaskParams) error {
	_, err := q.db.ExecContext(ctx, deleteTask, arg.ID, arg.UserID)
	return err
}

const getTask = `-- name: GetTask :one
SELECT id, user_id, title, created_at, updated_at FROM tasks17
WHERE id = $1 AND user_id = $2 LIMIT 1
`

type GetTaskParams struct {
	ID     int32
	UserID sql.NullInt32
}

func (q *Queries) GetTask(ctx context.Context, arg GetTaskParams) (Tasks17, error) {
	row := q.db.QueryRowContext(ctx, getTask, arg.ID, arg.UserID)
	var i Tasks17
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password FROM users17 WHERE email=$1
`

func (q *Queries) GetUser(ctx context.Context, email sql.NullString) (Users17, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i Users17
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, user_id, title, created_at, updated_at FROM tasks17
WHERE user_id = $1 ORDER BY id ASC
`

func (q *Queries) ListTasks(ctx context.Context, userID sql.NullInt32) ([]Tasks17, error) {
	rows, err := q.db.QueryContext(ctx, listTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tasks17
	for rows.Next() {
		var i Tasks17
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTask = `-- name: UpdateTask :one
UPDATE tasks17
  set title = $3,
  updated_at = $4
WHERE id = $1 AND user_id = $2
RETURNING id, user_id, title, created_at, updated_at
`

type UpdateTaskParams struct {
	ID        int32
	UserID    sql.NullInt32
	Title     sql.NullString
	UpdatedAt sql.NullString
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (Tasks17, error) {
	row := q.db.QueryRowContext(ctx, updateTask,
		arg.ID,
		arg.UserID,
		arg.Title,
		arg.UpdatedAt,
	)
	var i Tasks17
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
