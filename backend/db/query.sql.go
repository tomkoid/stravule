// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addFilter = `-- name: AddFilter :exec
INSERT INTO filters (user_id, filter_text, category)
VALUES (
  (SELECT id FROM users WHERE user_hash = $1),
  $2,
  $3
)
`

type AddFilterParams struct {
	UserHash   string
	FilterText string
	Category   pgtype.Text
}

func (q *Queries) AddFilter(ctx context.Context, arg AddFilterParams) error {
	_, err := q.db.Exec(ctx, addFilter, arg.UserHash, arg.FilterText, arg.Category)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    user_hash, sid
) VALUES (
    $1, $2
)
RETURNING id, user_hash, sid
`

type CreateUserParams struct {
	UserHash string
	Sid      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.UserHash, arg.Sid)
	var i User
	err := row.Scan(&i.ID, &i.UserHash, &i.Sid)
	return i, err
}

const deleteFilter = `-- name: DeleteFilter :exec
DELETE FROM filters
WHERE user_id = (SELECT id FROM users WHERE user_hash = $1)
  AND filter_text = $2
  AND category = $3
`

type DeleteFilterParams struct {
	UserHash   string
	FilterText string
	Category   pgtype.Text
}

func (q *Queries) DeleteFilter(ctx context.Context, arg DeleteFilterParams) error {
	_, err := q.db.Exec(ctx, deleteFilter, arg.UserHash, arg.FilterText, arg.Category)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, user_hash, sid FROM users
WHERE user_hash = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userHash string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userHash)
	var i User
	err := row.Scan(&i.ID, &i.UserHash, &i.Sid)
	return i, err
}

const listFilters = `-- name: ListFilters :many
SELECT id, user_id, filter_text, category, created_at FROM filters
WHERE (SELECT id FROM users WHERE user_hash = $1) = user_id
`

func (q *Queries) ListFilters(ctx context.Context, userHash string) ([]Filter, error) {
	rows, err := q.db.Query(ctx, listFilters, userHash)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Filter
	for rows.Next() {
		var i Filter
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FilterText,
			&i.Category,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
