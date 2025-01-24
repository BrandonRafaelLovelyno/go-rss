-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedByUserId :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetFollowedFeedsByUserId :many
SELECT f.* FROM feeds f JOIN feeds_follows ff ON f.id = ff.feed_id WHERE ff.user_id = $1;
