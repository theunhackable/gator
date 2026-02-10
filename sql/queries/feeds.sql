
-- name: CreateFeed :one
INSERT INTO feeds (user_id, name, url, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: ResetFeedTable :exec
DELETE FROM users;

