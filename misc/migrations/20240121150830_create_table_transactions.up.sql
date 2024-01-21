CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    wallet_id VARCHAR NOT NULL REFERENCES wallets(id),
    date BIGINT DEFAULT EXTRACT(epoch FROM NOW()) * 1000,
    type VARCHAR NOT NULL,
    gram FLOAT NOT NULL CHECK (gram >= 0.001),
    price_id VARCHAR NOT NULL REFERENCES prices(id),
    balance FLOAT NOT NULL
);

INSERT INTO transactions(wallet_id, type, gram, price_id, balance)
VALUES ('w001', 'topup', 0.01, 'p001', 0);