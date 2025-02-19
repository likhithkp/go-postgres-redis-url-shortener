CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_url TEXT UNIQUE NOT NULL,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);