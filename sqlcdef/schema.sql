CREATE TABLE users (
    id TEXT PRIMARY KEY,
    email TEXT NOT NULL
);

CREATE TABLE products (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGER NOT NULL
);

CREATE TABLE product_stocks (
    id TEXT PRIMARY KEY,
    product_id TEXT NOT NULL,
    stock_in INTEGER NOT NULL DEFAULT 0,
    stock_out INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE carts (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users (id)
);


CREATE TABLE cart_items (
    id TEXT PRIMARY KEY,
    cart_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL,
    
    FOREIGN KEY (cart_id) REFERENCES carts (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

CREATE TABLE orders (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE order_items (
    id TEXT PRIMARY KEY,
    order_id TEXT NOT NULL,
    product_id TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL,
    
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);
