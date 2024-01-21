CREATE TABLE IF NOT EXISTS prices (
    id VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    admin_id VARCHAR NOT NULL REFERENCES employees(id),
    topup INT NOT NULL,
    buyback INT NOT NULL,
    created_at BIGINT DEFAULT EXTRACT(epoch FROM NOW()) * 1000
);

INSERT INTO prices(id, admin_id, topup, buyback)
VALUES ('p001', 'a001', 70000, 60000);