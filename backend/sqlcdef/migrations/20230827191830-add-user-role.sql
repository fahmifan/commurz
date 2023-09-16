
-- +migrate Up
ALTER TABLE users ADD COLUMN IF NOT EXISTS "role" TEXT NOT NULL DEFAULT 'customer';

-- +migrate Down
