CREATE TABLE revoked_refresh_tokens (
    token TEXT PRIMARY KEY,
    revoked_at TIMESTAMP DEFAULT now()
);
