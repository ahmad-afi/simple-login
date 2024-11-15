CREATE TABLE users (
    ID VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    role VARCHAR(100),
    username VARCHAR(100),
    password VARCHAR(255),
    email VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
);

CREATE TABLE products (
    ID VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255),
    stok INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
);

CREATE TABLE history_stocks (
    ID VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) REFERENCES products(ID),
    created_by VARCHAR(36) REFERENCES users(ID),
    type ENUM('in', 'out'), -- hanya menerima 'in' atau 'out'
    description TEXT,
    total INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
);
