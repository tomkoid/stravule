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

-- name: ListFilters :many
SELECT * FROM filters
WHERE (SELECT id FROM users WHERE user_hash = $1) = user_id;

-- name: AddFilter :exec
INSERT INTO filters (user_id, filter_text, category)
VALUES (
  (SELECT id FROM users WHERE user_hash = $1),
  $2,
  $3
);

-- name: DeleteFilter :exec
DELETE FROM filters
WHERE user_id = (SELECT id FROM users WHERE user_hash = $1)
  AND filter_text = $2
  AND category = $3;
