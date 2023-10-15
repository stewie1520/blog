// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: post.query.sql

package dao_post

import (
	"context"

	"github.com/google/uuid"
)

const countByUserId = `-- name: CountByUserId :one
SELECT COUNT(*) FROM "posts" WHERE "posts"."user_id" = $1
`

func (q *Queries) CountByUserId(ctx context.Context, userID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countByUserId, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const create = `-- name: Create :one
INSERT INTO "posts" ("id", "content", "user_id") VALUES ($1, $2, $3) RETURNING id, content, user_id, created_at, updated_at, deleted_at
`

type CreateParams struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"user_id"`
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (Post, error) {
	row := q.db.QueryRow(ctx, create, arg.ID, arg.Content, arg.UserID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const findById = `-- name: FindById :one
SELECT id, content, user_id, created_at, updated_at, deleted_at FROM "posts" WHERE "posts"."id" = $1 LIMIT 1
`

func (q *Queries) FindById(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, findById, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listByUserId = `-- name: ListByUserId :many
SELECT id, content, user_id, created_at, updated_at, deleted_at FROM "posts" WHERE "posts"."user_id" = $1 ORDER BY "posts"."created_at" DESC LIMIT $2 OFFSET $3
`

type ListByUserIdParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListByUserId(ctx context.Context, arg ListByUserIdParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, listByUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const removePost = `-- name: RemovePost :exec
UPDATE "posts" SET "deleted_at" = now() WHERE "id" = $1
`

func (q *Queries) RemovePost(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, removePost, id)
	return err
}

const updatePost = `-- name: UpdatePost :exec
UPDATE "posts" SET "content" = $2, "updated_at" = now() WHERE "id" = $1
`

type UpdatePostParams struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.Exec(ctx, updatePost, arg.ID, arg.Content)
	return err
}