--------------------------------------------------
-- Query
--------------------------------------------------

-- name: FindAllCartItemsByCartIDs :many
SELECT * FROM cart_items WHERE cart_id = ANY(@cart_ids::TEXT[]);

-- name: FindOrderItemsByOrderID :many
SELECT * FROM order_items WHERE order_id = @order_id;

-- name: FindOrderByID :one
SELECT * FROM orders WHERE id = @id;

-------------------------------------------------- 
-- Mutation
--------------------------------------------------

-- name: SaveCartItem :one
INSERT INTO cart_items (id, cart_id, product_id, quantity, price) 
VALUES (@id, @cart_id, @product_id, @quantity, @price)
ON CONFLICT (id) DO UPDATE
    SET quantity = @quantity, price = @price
RETURNING *;

-- name: CreateCart :one
INSERT INTO carts (id, user_id)
VALUES (@id, @user_id)
RETURNING *;

-- name: SaveOrder :one
INSERT INTO orders (id, user_id, number)
VALUES (@id, @user_id, @number)
RETURNING *;

-- name: SaveOrderItem :one
INSERT INTO order_items (id, order_id, product_id, quantity, price)
VALUES (@id, @order_id, @product_id, @quantity, @price)
RETURNING *;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = @id;

-- name: DeleteAllCartItem :exec
DELETE FROM cart_items WHERE id = ANY(@id::TEXT[]);