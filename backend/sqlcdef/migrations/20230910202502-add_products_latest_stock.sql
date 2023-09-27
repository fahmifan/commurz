
-- +migrate Up
ALTER TABLE products ADD COLUMN IF NOT EXISTS latest_stock BIGINT NOT NULL DEFAULT 0;
-- +migrate Down

ALTER TABLE products DROP COLUMN IF EXISTS latest_stock;

