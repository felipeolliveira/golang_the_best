-- Write your migrate up statements here
--
CREATE TABLE business_account (
    -- CUID must be generate on server
    id VARCHAR(30) PRIMARY KEY,
    revenue DECIMAL(18,2),
    age INT,
    trade_name VARCHAR(255),
    phone VARCHAR(20),
    email VARCHAR(255),
    category VARCHAR(50),
    balance DECIMAL(18,2),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    closed_at TIMESTAMPTZ,

    CONSTRAINT unique_business_account_email UNIQUE (email)
);

CREATE INDEX idx_business_account_phone ON business_account(phone);
CREATE INDEX idx_business_account_category ON business_account(category);

---- create above / drop below ----

DROP INDEX IF EXISTS idx_business_account_phone;
DROP INDEX IF EXISTS idx_business_account_category;
DROP TABLE IF EXISTS business_account;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
