CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT UNIQUE,
    avatar TEXT UNIQUE,
    classic_points NUMERIC(7, 4) DEFAULT 0,
    platformer_points NUMERIC(7, 4) DEFAULT 0,
    discord TEXT UNIQUE,
    youtube TEXT UNIQUE,
    twitter TEXT UNIQUE,
    twitch TEXT UNIQUE,
    is_flagged BOOLEAN DEFAULT false
);
