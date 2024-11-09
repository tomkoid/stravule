-- name: GetUser :one
SELECT * FROM users
WHERE userHash = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    userHash, sid
) VALUES (
    $1, $2
)
RETURNING *;
