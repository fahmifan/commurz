CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  verify_token TEXT NOT NULL,
  "status" VARCHAR(64) NOT NULL,
  last_login_at TIMESTAMP,
  archived boolean NOT NULL DEFAULT false,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    price BIGINT NOT NULL,
    version BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS product_stocks (
    id TEXT PRIMARY KEY,
    product_id TEXT NOT NULL,
    stock_in BIGINT NOT NULL DEFAULT 0,
    stock_out BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE IF NOT EXISTS carts (
    id TEXT PRIMARY KEY,
    user_id uuid NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE IF NOT EXISTS cart_items (
    id TEXT PRIMARY KEY,
    cart_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    quantity BIGINT NOT NULL,
    price BIGINT NOT NULL,
    
    FOREIGN KEY (cart_id) REFERENCES carts (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE IF NOT EXISTS orders (
    id TEXT PRIMARY KEY,
    user_id uuid NOT NULL,
    number TEXT NOT NULL,
    
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS order_items (
    id TEXT PRIMARY KEY,
    order_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    quantity BIGINT NOT NULL,
    price BIGINT NOT NULL,

    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);
