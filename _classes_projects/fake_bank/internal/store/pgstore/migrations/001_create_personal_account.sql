-- Write your migrate up statements here

CREATE TABLE personal_account (
    -- CUID must be generate on server
    id CHAR(30) PRIMARY KEY,
    monthy_income DECIMAL(15,2) NOT NULL DEFAULT 0,
    age INTEGER NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255) NOT NULL,
    category VARCHAR(50),
    balance DECIMAL(15,2) NOT NULL DEFAULT 0,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    closed_at TIMESTAMPTZ,

    CONSTRAINT unique_personal_account_email UNIQUE (email)
);

CREATE INDEX idx_personal_account_email ON personal_account(email);
CREATE INDEX idx_personal_account_phone ON personal_account(phone);
CREATE INDEX idx_personal_account_category ON personal_account(category);

---- create above / drop below ----

DROP INDEX IF EXISTS idx_personal_account_category;
DROP INDEX IF EXISTS idx_personal_account_phone;
DROP INDEX IF EXISTS idx_personal_account_email;
DROP TABLE IF EXISTS personal_account;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
