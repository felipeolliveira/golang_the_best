-- name: CreateBusinessAccount :one
INSERT INTO business_account (id, email, trade_name, age, phone, category)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetBusinessAccount :one
SELECT * FROM business_account
WHERE id = $1;

-- name: UpdateBusinessAccount :one
UPDATE business_account
SET email = $2,
    trade_name = $3,
    age = $4,
    phone = $5,
    category = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: IncrementBalanceBusinessAccount :exec
UPDATE business_account
SET balance = balance + @value::numeric
WHERE id = $1;

-- name: DecrementBalanceBusinessAccount :exec
UPDATE business_account
SET balance = balance - @value::numeric
WHERE id = $1;

-- name: CloseBusinessAccount :exec
UPDATE business_account
SET closed_at = now();

-- name: ReopenBusinessAccount :one
UPDATE business_account
SET closed_at = NULL
RETURNING *;
