
-- +migrate Up
ALTER TABLE products ADD COLUMN latest_stock BIGINT NOT NULL DEFAULT 0;
-- +migrate Down

ALTER TABLE products DROP COLUMN latest_stock;

