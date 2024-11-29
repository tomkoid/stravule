-- name: GetUser :one
SELECT id, user_hash, sid, order_days_exceptions, is_beta_tester FROM users
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

-- name: RegisterBetatester :exec
UPDATE users
SET is_beta_tester = true
WHERE user_hash = $1;

-- name: AddWeekdayOrderingException :exec
UPDATE users
SET order_days_exceptions = array_append(order_days_exceptions, $2)
WHERE user_hash = $1;


-- name: RemoveWeekdayOrderingException :exec
UPDATE users
SET order_days_exceptions = array_remove(order_days_exceptions, $2)
WHERE user_hash = $1;
