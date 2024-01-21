CREATE TABLE IF NOT EXISTS wallets (
    id VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    customer_id VARCHAR NOT NULL REFERENCES customers(id),
    balance FLOAT DEFAULT 0
);

INSERT INTO wallets(id, customer_id)
VALUES ('w001', 'c001');