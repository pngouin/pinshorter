CREATE TABLE IF NOT EXISTS user (
    user_id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TEXT NOT NULL,
    deleted_at TEXT
);

CREATE TABLE IF NOT EXISTS link (
    link_id TEXT PRIMARY KEY ,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    api_point TEXT NOT NULL,
    created_at TEXT NOT NULL,
    deleted_at TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

CREATE TABLE IF NOT EXISTS count (
    count_id TEXT PRIMARY KEY,
    date TEXT NOT NULL,
    link_id TEXT NOT NULL,
    FOREIGN KEY (link_id) REFERENCES link(link_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS params (
    salt TEXT
)
