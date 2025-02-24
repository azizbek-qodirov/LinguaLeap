CREATE TABLE user_lessons (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    lesson_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_data (
    user_id UUID PRIMARY KEY,
    level VARCHAR(100),
    native_lang VARCHAR(100),
    xp BIGINT DEFAULT 0,
    daily_streak INT DEFAULT 0,
    played_games_count BIGINT DEFAULT 0,
    winning_percentage FLOAT DEFAULT 0
);