-- name: FindUserById :one
SELECT * FROM "users" WHERE "users"."id" = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO "users" ("id", "full_name", "account_id", "bio") VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByAccountID :one
SELECT * FROM "users" WHERE "users"."account_id" = $1 LIMIT 1;