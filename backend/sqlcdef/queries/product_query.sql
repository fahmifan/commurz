-- name: FindAllProductsForBackoffice :many
SELECT * FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END
LIMIT @page_limit
OFFSET @page_offset;

-- name: CountAllProductsForBackoffice :one
SELECT COUNT(*) FROM products
WHERE 
    CASE WHEN @set_name::bool THEN ("name" LIKE '%' || @name || '%') ELSE TRUE END;