-- USERS TABLE --
CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    username VARCHAR (100) UNIQUE NOT NULL,
    password_hash VARCHAR (100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    changed_at TIMESTAMP
    currency VARCHAR(10)
);

CREATE OR REPLACE FUNCTION update_users_changed_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.changed_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER user_update_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_users_changed_at();

INSERT INTO users(username, password) VALUES('Sbeve', 'SbevePass');

UPDATE users
SET username = 'Sbeve'
WHERE username = '_Sbeve';

SELECT * FROM users;

-- STOCKS TABLE --

DROP TABLE IF EXISTS stocks;
CREATE TABLE stocks(
    id serial PRIMARY KEY,
    ticker VARCHAR (100) UNIQUE NOT NULL,
    prev_close REAL,
    current_price REAL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
    currency VARCHAR(10)
);

CREATE OR REPLACE FUNCTION update_stock_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER stock_update_trigger
BEFORE UPDATE ON stocks
FOR EACH ROW
EXECUTE FUNCTION update_stock_updated_at();

INSERT INTO stocks(ticker, prev_close, current_price) 
VALUES('AAPL', 153.2, 154.1);

UPDATE stocks
SET current_price = 55.1
WHERE ticker='AAPL'; 
SELECT * FROM stocks;

-- PORTFOLIO TABLE -- 
DROP TABLE IF EXISTS stocks_in_portfolio;
CREATE TABLE stocks_in_portfolio(
    user_id INT REFERENCES users(id) NOT NULL,
    stock_id INT REFERENCES stocks(id) NOT NULL,
    shares INT NOT NULL,
    CONSTRAINT unique_user_stock_pair UNIQUE (user_id, stock_id)  -- Ensures one user can't own the same stock multiple times
);

WITH user_stock_ids AS (
  SELECT
    (SELECT id FROM users WHERE username = 'Sbeve') AS user_id,
    (SELECT id FROM stocks WHERE ticker = 'AAPL') AS stock_id
)
INSERT INTO stocks_in_portfolio (user_id, stock_id, shares)
VALUES (
    (SELECT user_id FROM user_stock_ids), 
    (SELECT stock_id FROM user_stock_ids), 
    1
    );

SELECT
    s.ticker AS stock_ticker,
    s.current_price AS current_price,
    sp.shares,
    s.current_price * sp.shares AS total
FROM stocks_in_portfolio sp
JOIN users u ON sp.user_id = u.id
JOIN stocks s ON sp.stock_id = s.id
WHERE u.username = 'Sbeve';

SELECT
    s.ticker AS stock_ticker,
    s.current_price AS current_price,
    s.currency AS currency,
    sp.shares,
    s.current_price * sp.shares AS total
FROM
    stocks_in_portfolio sp
JOIN
    users u ON sp.user_id = u.id
JOIN
    stocks s ON sp.stock_id = s.id
WHERE
    u.username = 'Sbeve';

-- CONVERSION RATES --
DROP TABLE IF EXISTS conversion_rates;
CREATE TABLE conversion_rates(
    id serial PRIMARY KEY,
    currency_pair VARCHAR(20),
    conversion_rate REAL,
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL
)

CREATE OR REPLACE FUNCTION update_conversion_rates_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER conversion_rates_update_trigger
BEFORE UPDATE ON conversion_rates
FOR EACH ROW
EXECUTE FUNCTION update_stock_updated_at();