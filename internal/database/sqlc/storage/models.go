// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

import (
	"database/sql"
)

type Tasks17 struct {
	ID        int32
	UserID    sql.NullInt32
	Title     sql.NullString
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}

type Users17 struct {
	ID       int32
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
}
