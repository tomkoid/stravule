-- name: GetUser :one
SELECT * FROM users
WHERE user_hash = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    user_hash, sid
) VALUES (
    $1, $2
)
RETURNING *;

-- name: AddFilter :exec
INSERT INTO filters (user_id, filter_text, category)
VALUES (
  (SELECT id FROM users WHERE user_hash = $1),
  $2,
  $3
);
