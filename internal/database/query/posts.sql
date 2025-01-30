-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, published_at, title, description, url, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserPosts :many
SELECT posts.* FROM posts 
JOIN feeds ON post.feed_id = feeds.id
JOIN feeds_follows ON feeds_follows.feed_id = feeds.id
WHERE feeds_follows.user_id = $1;
