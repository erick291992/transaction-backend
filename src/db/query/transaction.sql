-- name: CreateTransaction :one
INSERT INTO transactions (
	account_id,
	operation_type,
	amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;