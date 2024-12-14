-- name: GetUser :one
SELECT id, user_hash, sid, order_days_exceptions, no_order_days, is_beta_tester FROM users
WHERE user_hash = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    user_hash, sid
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateUserSID :exec
UPDATE users
SET sid = $2
WHERE user_hash = $1;

-- name: ListFilters :many
SELECT * FROM filters
WHERE (SELECT id FROM users WHERE user_hash = $1) = user_id;

-- name: AddFilter :exec
INSERT INTO filters (user_id, filter_text, category, weight)
VALUES (
  (SELECT id FROM users WHERE user_hash = $1),
  $2,
  $3,
  $4
);

-- name: DeleteFilter :exec
DELETE FROM filters
WHERE user_id = (SELECT id FROM users WHERE user_hash = $1)
  AND filter_text = $2
  AND category = $3; 

-- name: UpdateFilterWeight :exec
UPDATE filters
SET weight = $2
WHERE user_id = (SELECT id FROM users WHERE user_hash = $1)
  AND filter_text = $3
  AND category = $4;

-- name: RegisterBetatester :exec
UPDATE users
SET is_beta_tester = true
WHERE user_hash = $1;

-- name: ListWeekdayOrderingExceptions :exec
SELECT order_days_exceptions FROM users
WHERE user_hash = $1;

-- name: ListNoOrderDays :exec
SELECT no_order_days FROM users
WHERE user_hash = $1;

-- name: AddWeekdayOrderingException :exec
UPDATE users
SET order_days_exceptions = array_append(order_days_exceptions, $2)
WHERE user_hash = $1;

-- name: AddNoOrderDay :exec
UPDATE users
SET no_order_days = array_append(no_order_days, $2)
WHERE user_hash = $1;

-- name: RemoveWeekdayOrderingException :exec
UPDATE users
SET order_days_exceptions = array_remove(order_days_exceptions, $2)
WHERE user_hash = $1;

-- name: RemoveNoOrderDay :exec
UPDATE users
SET no_order_days = array_remove(no_order_days, $2)
WHERE user_hash = $1;
