CREATE TABLE IF NOT EXISTS customers (
    id VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL
);

INSERT INTO customers(id, username, password, role)
VALUES ('c001', 'john', 'doe', 'customer');