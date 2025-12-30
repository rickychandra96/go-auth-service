-- Create users table
CREATE TABLE IF NOT EXISTS users
(
    id         UUID                NOT NULL PRIMARY KEY,
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    name       VARCHAR(255)        NOT NULL,
    is_active  BOOLEAN             NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP           NULL
);

-- Create index on email for faster lookups
CREATE INDEX idx_users_email ON users (email) WHERE deleted_at IS NULL;

-- Create index on deleted_at for soft delete queries
CREATE INDEX idx_users_deleted_at ON users (deleted_at);