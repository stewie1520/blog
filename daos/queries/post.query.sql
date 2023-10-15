-- name: Create :one
INSERT INTO "posts" ("id", "content", "user_id") VALUES ($1, $2, $3) RETURNING *;

-- name: ListByUserId :many
SELECT * FROM "posts" WHERE "posts"."user_id" = $1 AND "posts"."deleted_at" IS NULL ORDER BY "posts"."created_at" DESC LIMIT $2 OFFSET $3;

-- name: CountByUserId :one
SELECT COUNT(*) FROM "posts" WHERE "posts"."user_id" = $1 AND "posts"."deleted_at" IS NULL;

-- name: FindById :one
SELECT * FROM "posts" WHERE "posts"."id" = $1 AND "posts"."deleted_at" IS NULL LIMIT 1;

-- name: UpdatePost :exec
UPDATE "posts" SET "content" = $2, "updated_at" = now() WHERE "id" = $1 AND "posts"."deleted_at" IS NULL;

-- name: RemovePost :exec
UPDATE "posts" SET "deleted_at" = now() WHERE "id" = $1 AND "posts"."deleted_at" IS NULL;
