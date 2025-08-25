-- +goose Up
CREATE TABLE feeds(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT ,
    URL TEXT UNIQUE,
    userid UUID NOT NULL REFERENCES users
    ON DELETE CASCADE,
    FOREIGN KEY(userid) REFERENCES users(id)
); 

-- +goose Down
DROP TABLE feeds; 