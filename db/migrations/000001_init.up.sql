-- Create the Accounts table
CREATE TABLE Accounts (
    id SERIAL PRIMARY KEY,
    owner VARCHAR(255) NOT NULL,
    balance DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on owner for faster retrieval
CREATE INDEX idx_owner ON Accounts(owner);

-- Create the Entries table
CREATE TABLE Entries (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES Accounts(id) ON DELETE CASCADE,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on account_id for faster retrieval
CREATE INDEX idx_account_id ON Entries(account_id);

-- Create the Transfers table
CREATE TABLE Transfers (
    id SERIAL PRIMARY KEY,
    from_account_id INTEGER REFERENCES Accounts(id) ON DELETE CASCADE,
    to_account_id INTEGER REFERENCES Accounts(id) ON DELETE CASCADE,
    amount DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on from_account_id and to_account_id for faster retrieval
CREATE INDEX idx_transfer_accounts ON Transfers(from_account_id, to_account_id);
