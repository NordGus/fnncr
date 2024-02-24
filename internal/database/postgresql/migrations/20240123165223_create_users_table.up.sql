CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY NOT NULL,
    username VARCHAR(20) UNIQUE NOT NULL,
    password_digest VARCHAR NOT NULL,
    session_version INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);