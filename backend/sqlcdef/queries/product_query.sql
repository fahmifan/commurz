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
