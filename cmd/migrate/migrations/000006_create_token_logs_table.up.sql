CREATE TABLE IF NOT EXISTS token_logs (
    id SERIAL PRIMARY KEY,
    client_token TEXT NOT NULL REFERENCES clients(client_token) ON DELETE CASCADE,
    token_type TEXT NOT NULL,          
    action TEXT NOT NULL,              
    token TEXT,                         
    ip_address TEXT,
    user_agent TEXT,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT now()
);
