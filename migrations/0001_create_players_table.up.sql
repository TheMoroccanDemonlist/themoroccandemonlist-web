CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    username TEXT UNIQUE NOT NULL,
    avatar TEXT UNIQUE NOT NULL,
    classic_points NUMERIC(7, 4) DEFAULT 0,
    platformer_points NUMERIC(7, 4) DEFAULT 0,
    discord TEXT UNIQUE NOT NULL,
    youtube TEXT UNIQUE NOT NULL,
    twitter TEXT UNIQUE NOT NULL,
    twitch TEXT UNIQUE NOT NULL,
    is_flagged BOOLEAN DEFAULT false
);
