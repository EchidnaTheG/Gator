-- +goose Up

CREATE TABLE feed_follows(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    userid UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feedid UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    CONSTRAINT unique_userid_feedid   UNIQUE(userid, feedid)
); 

-- +goose Down
DROP TABLE feed_follows; 