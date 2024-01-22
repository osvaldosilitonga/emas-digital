CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    wallet_id VARCHAR NOT NULL REFERENCES wallets(id),
    date BIGINT DEFAULT EXTRACT(epoch FROM NOW()) * 1000,
    type VARCHAR NOT NULL,
    gram FLOAT(3) NOT NULL CHECK (gram >= 0.001),
    price_id VARCHAR NOT NULL REFERENCES prices(id),
    balance FLOAT(3) NOT NULL
);

INSERT INTO transactions(wallet_id, type, gram, price_id, balance, date)
VALUES 
    ('w001', 'topup', 0.2, 'p001', 0.7, 1705510800000),
    ('w001', 'buyback', 0.1, 'p001', 0.6, 1705597200000),
    ('w001', 'buyback', 0.1, 'p001', 0.5, 1705683600000),
    ('w001', 'topup', 0.2, 'p001', 0.7, 1705770000000),
    ('w001', 'topup', 0.1, 'p001', 0.8, 1705856400000);