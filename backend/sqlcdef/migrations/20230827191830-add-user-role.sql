
-- +migrate Up
ALTER TABLE users ADD COLUMN "role" TEXT NOT NULL DEFAULT 'customer';

-- +migrate Down
