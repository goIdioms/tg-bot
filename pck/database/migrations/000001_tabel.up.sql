CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    answer TEXT NOT NULL,
    level VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS user_states (
    chat_id BIGINT PRIMARY KEY,
    step INTEGER NOT NULL,
    task_question TEXT,
    task_answer TEXT,
    task_level VARCHAR(16),
    message_id INTEGER
);