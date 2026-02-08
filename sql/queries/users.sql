-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT name FROM users
WHERE name = $1;

-- name: ResetUserTable :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
