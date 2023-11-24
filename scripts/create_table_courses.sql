CREATE TABLE courses (
	id INTEGER primary key autoincrement,
	title TEXT,
	description TEXT,
    created_at TEXT, -- ISO8601: YYYY-MM-DD HH:MM:SS.SSS
    updated_at TEXT,
    deleted_at TEXT
);
