CREATE TABLE IF NOT EXISTS users
(
    id       INTEGER PRIMARY KEY NOT NULL,
    name     TEXT                NOT NULL,
    email    TEXT            UNIQUE    NOT NULL,
    phone    TEXT        UNIQUE        NOT NULL,
    password TEXT                NOT NULL
);

CREATE TABLE IF NOT EXISTS posts
(
    id      INTEGER PRIMARY KEY NOT NULL,
    user_id INTEGER             NOT NULL,
    post    TEXT                NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users (name, email, phone, password) VALUES ('David', 'david.christiea@gmail.com', '1234567890', 'password');