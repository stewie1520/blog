-- name: CreateAccount :one
INSERT INTO "accounts" ("id", "email", "password") VALUES ($1, $2, $3) RETURNING *;

-- name: FindByEmail :one
SELECT * FROM "accounts" WHERE "email" = $1 LIMIT 1;