-- name: FindOrderItemsByOrderID :many
SELECT * FROM order_items WHERE order_id = ?;

-- name: FindOrderByID :one
SELECT * FROM orders WHERE id = ?;

-- name: SaveOrder :one
INSERT INTO orders (id, user_id, number)
VALUES (@id, @user_id, @number)
RETURNING *;

-- name: SaveOrderItem :one
INSERT INTO order_items (id, order_id, product_id, quantity, price)
VALUES (@id, @order_id, @product_id, @quantity, @price)
RETURNING *;
