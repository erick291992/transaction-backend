-- name: CreateOperationsType :one
INSERT INTO operations_types (
	"description"
) VALUES (
  $1
)
RETURNING *;

-- name: GetOperationsType :one
SELECT * FROM operations_types
WHERE id = $1 LIMIT 1;

-- name: ListOperationsTypes :many
SELECT * FROM operations_types
ORDER BY id
LIMIT $1
OFFSET $2;