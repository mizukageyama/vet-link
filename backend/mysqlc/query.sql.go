// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package mysqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (given_name, family_name, email, password) VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	GivenName  string
	FamilyName string
	Email      string
	Password   string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.GivenName,
		arg.FamilyName,
		arg.Email,
		arg.Password,
	)
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, given_name, family_name, email, password FROM users WHERE email = ? LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.GivenName,
		&i.FamilyName,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, given_name, family_name, email, password FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.GivenName,
			&i.FamilyName,
			&i.Email,
			&i.Password,
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
