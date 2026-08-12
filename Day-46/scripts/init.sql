-- Initialize Schema for Day 46 Go API Microservice

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Seed Initial Data
INSERT INTO users (name, email, role) VALUES
    ('Golang Developer', 'go.dev@example.com', 'admin'),
    ('Docker Specialist', 'docker@example.com', 'engineer'),
    ('DevOps Architect', 'devops@example.com', 'lead')
ON CONFLICT (email) DO NOTHING;
