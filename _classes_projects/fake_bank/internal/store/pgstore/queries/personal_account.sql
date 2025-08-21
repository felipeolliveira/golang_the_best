-- name: CreatePersonalAccount :one
INSERT INTO personal_account (id, email, full_name, age, phone, category)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPersonalAccount :one
SELECT * FROM personal_account
WHERE id = $1;

-- name: UpdatePersonalAccount :one
UPDATE personal_account
SET email = $2,
    full_name = $3,
    age = $4,
    phone = $5,
    category = $6,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: IncrementBalancePersonalAccount :exec
UPDATE personal_account
SET balance = balance + @value::numeric
WHERE id = $1;

-- name: DecrementBalancePersonalAccount :exec
UPDATE personal_account
SET balance = balance - @value::numeric
WHERE id = $1;

-- name: ClosePersonalAccount :exec
UPDATE personal_account
SET closed_at = now();

-- name: ReopenPersonalAccount :one
UPDATE personal_account
SET closed_at = NULL
RETURNING *;

