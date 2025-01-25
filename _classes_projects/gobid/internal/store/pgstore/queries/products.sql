-- name: GetProductById :one
SELECT
  id, seller_id, product_name,
  description, base_price_in_cents,
  auction_end, is_sold,
  created_at, updated_at
FROM products
WHERE id = $1

