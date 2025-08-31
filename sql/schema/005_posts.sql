-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    URL TEXT UNIQUE,
    description TEXT,
    published_at TIMESTAMP,
    feedid UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
);

-- +goose Down
DROP TABLE posts;