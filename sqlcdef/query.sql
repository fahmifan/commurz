-- name: FindProductByID :one
SELECT * FROM products WHERE id = ?;

-- name: FindCartByUserID :one
SELECT * FROM carts WHERE user_id = ?;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: FindAllCartItemsByCartIDs :many
SELECT * FROM cart_items WHERE cart_id IN (sqlc.slice('cart_ids'));

-- name: FindAllProductsByIDs :many
SELECT * FROM products WHERE id IN (sqlc.slice('product_ids'));

-- name: FindAllProductStocksByIDs :many
SELECT * FROM product_stocks WHERE product_id IN (sqlc.slice('product_ids'));

-- name: SaveCartItem :one
INSERT INTO cart_items (id, cart_id, product_id, quantity, price) 
VALUES (@id, @cart_id, @product_id, @quantity, @price)
RETURNING *;

-- name: SaveOrder :one
INSERT INTO orders (id, user_id)
VALUES (@id, @user_id)
RETURNING *;

-- name: SaveOrderItems :one
INSERT INTO order_items (id, order_id, product_id, quantity, price)
VALUES (@id, @order_id, @product_id, @quantity, @price)
RETURNING *;

-- name: SaveProduct :one
INSERT INTO products (id, name, price)
VALUES (@id, @name, @price)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET 
    name = @name, 
    price = @price 
WHERE 
    id = @id 
RETURNING *;

-- name: CreateProductStock :one
INSERT INTO product_stocks (id, product_id, stock_in, stock_out)
VALUES (@id, @product_id, @stock_in, @stock_out)
RETURNING *;

-- name: CreateCart :one
INSERT INTO carts (id, user_id)
VALUES (@id, @user_id)
RETURNING *;

-- name: CreateUser :one
INSERT INTO users (id, email)
VALUES (@id, @email)
RETURNING *;
