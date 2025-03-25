CREATE TABLE revoked_tokens (
    client_token TEXT PRIMARY KEY,
    revoked_at TIMESTAMP DEFAULT now()
);
