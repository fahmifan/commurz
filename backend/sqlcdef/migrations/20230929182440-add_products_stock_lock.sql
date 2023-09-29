
-- +migrate Up
CREATE TABLE product_stock_lock (
    id TEXT PRIMARY KEY,
    product_id TEXT NOT NULL,
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);

-- +migrate Down
