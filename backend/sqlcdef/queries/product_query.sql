-- name: FindProductByID :one
SELECT * FROM products WHERE id = @id;

-- name: FindAllProductsByIDs :many
SELECT * FROM products WHERE id = ANY(@product_ids::TEXT[]);

-- name: FindAllProductStocksByIDs :many
SELECT * FROM product_stocks WHERE product_id = ANY(@product_ids::TEXT[]);

-- name: FindAllProductsForBackoffice :many
SELECT * FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END
ORDER BY id DESC
LIMIT @page_limit
OFFSET @page_offset;

-- name: CountAllProductsForBackoffice :one
SELECT COUNT(*) FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END;

-- name: UpdateProduct :one
UPDATE products SET 
    name = @name, 
    price = @price,
    latest_stock = @latest_stock
WHERE 
    id = @id 
RETURNING *;

-- name: SaveProduct :one
INSERT INTO products (id, name, price)
VALUES (@id, @name, @price)
RETURNING *;

-- name: BumpProductVersion :one
UPDATE products SET version = version + 1 
WHERE id = @id AND version = @current_version
RETURNING *;

-- name: CreateProductStock :one
INSERT INTO product_stocks (id, product_id, stock_in, stock_out)
VALUES (@id, @product_id, @stock_in, @stock_out)
RETURNING *;

-- --------------------------------------------------
-- App & Backoffice is seperated since they will 
-- diverge later
-- --------------------------------------------------

-- name: FindAllProductsForApp :many
SELECT * FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END
ORDER BY id DESC
LIMIT @page_limit
OFFSET @page_offset;

-- name: CountAllProductsForApp :one
SELECT COUNT(*) FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END;
