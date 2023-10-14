// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.query.sql

package dao_user

import (
	"context"

	"github.com/google/uuid"
	"github.com/stewie1520/blog/tools/types"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "users" ("id", "full_name", "account_id", "bio") VALUES ($1, $2, $3, $4) RETURNING id, account_id, full_name, bio, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	ID        uuid.UUID        `json:"id"`
	FullName  string           `json:"full_name"`
	AccountID uuid.UUID        `json:"account_id"`
	Bio       types.NullString `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FullName,
		arg.AccountID,
		arg.Bio,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FullName,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findUserById = `-- name: FindUserById :one
SELECT id, account_id, full_name, bio, created_at, updated_at, deleted_at FROM "users" WHERE "users"."id" = $1 LIMIT 1
`

func (q *Queries) FindUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, findUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FullName,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByAccountID = `-- name: GetUserByAccountID :one
SELECT id, account_id, full_name, bio, created_at, updated_at, deleted_at FROM "users" WHERE "users"."account_id" = $1 LIMIT 1
`

func (q *Queries) GetUserByAccountID(ctx context.Context, accountID uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByAccountID, accountID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.FullName,
		&i.Bio,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
