-- name: GetFoodsForUserId :many
SELECT * FROM foodRecord
WHERE userId = $1;