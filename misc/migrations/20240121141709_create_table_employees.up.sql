CREATE TABLE IF NOT EXISTS employees (
    id VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    username VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL
);

INSERT INTO employees(id, username, password, role)
VALUES ('a001', 'admin', 'admin', 'admin');