
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

-- name: GetFeedDetails :many

SELECT u.name AS username, f.name AS feed_name, f.url  AS url
FROM users as u INNER JOIN feeds as f
ON f.user_id = u.id;
