-- name: Create :one
INSERT INTO "posts" ("id", "content", "user_id") VALUES ($1, $2, $3) RETURNING *;

-- name: ListByUserId :many
SELECT * FROM "posts" WHERE "posts"."user_id" = $1 ORDER BY "posts"."created_at" DESC LIMIT $2 OFFSET $3;

-- name: CountByUserId :one
SELECT COUNT(*) FROM "posts" WHERE "posts"."user_id" = $1;
