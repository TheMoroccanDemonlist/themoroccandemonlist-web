CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID REFERENCES players(id),
    email TEXT UNIQUE NOT NULL,
    sub TEXT UNIQUE NOT NULL,
    is_banned BOOLEAN DEFAULT false,
    is_deleted BOOLEAN DEFAULT false
);
