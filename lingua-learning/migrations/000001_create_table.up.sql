CREATE TABLE lessons (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    lang_1 VARCHAR(10) NOT NULL,
    lang_2 VARCHAR(10) NOT NULL,
    level VARCHAR(100) NOT NULL,
    order_number INT UNIQUE NOT NULL
);

-- CREATE TABLE IF NOT EXISTS users (
--     id UUID PRIMARY KEY,
--     username VARCHAR(255) NOT NULL UNIQUE,
--     email VARCHAR(255) NOT NULL UNIQUE,
--     password VARCHAR(255) NOT NULL
-- );

CREATE TABLE vocabularies (
    id UUID PRIMARY KEY,
    lesson_id UUID REFERENCES lessons(id),
    type VARCHAR(255) NOT NULL,
    question TEXT NOT NULL,
    options TEXT NOT NULL,
    correct_answer TEXT NOT NULL
);