-- Create refresh_tokens table
CREATE TABLE refresh_tokens
(
    id         UUID      NOT NULL PRIMARY KEY,
    user_id    UUID      NOT NULL,
    token      TEXT      NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_revoked BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create index on token for faster lookups
CREATE INDEX idx_refresh_tokens_token ON refresh_tokens (token);

-- Create index on user_id
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens (user_id);

-- Create index for cleanup of expired tokens
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens (expires_at);