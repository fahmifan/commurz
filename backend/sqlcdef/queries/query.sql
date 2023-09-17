-- name: FindProductByID :one
SELECT * FROM products WHERE id = @id;

-- name: FindCartByUserID :one
SELECT * FROM carts WHERE user_id = @user_id;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = @id;

-- name: FindAllCartItemsByCartIDs :many
SELECT * FROM cart_items WHERE cart_id = ANY(@cart_ids::TEXT[]);

-- name: FindAllProductsByIDs :many
SELECT * FROM products WHERE id = ANY(@product_ids::TEXT[]);

-- name: FindAllProductStocksByIDs :many
SELECT * FROM product_stocks WHERE product_id = ANY(@product_ids::TEXT[]);

-- name: SaveCartItem :one
INSERT INTO cart_items (id, cart_id, product_id, quantity, price) 
VALUES (@id, @cart_id, @product_id, @quantity, @price)
RETURNING *;

-- name: SaveProduct :one
INSERT INTO products (id, name, price)
VALUES (@id, @name, @price)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET 
    name = @name, 
    price = @price,
    latest_stock = @latest_stock
WHERE 
    id = @id 
RETURNING *;

-- name: BumpProductVersion :one
UPDATE products SET version = version + 1 
WHERE id = @id AND version = @current_version
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

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = @email LIMIT 1;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = @id;