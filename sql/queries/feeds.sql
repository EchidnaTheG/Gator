-- name: CreateFeed :one
INSERT INTO feeds ( id, created_at, updated_at,name , URL, userid)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetFeeds :many
SELECT * 
FROM feeds
ORDER BY id;

-- name: LookUpFeedByURL :one

SELECT *
FROM FEEDS
WHERE URL = $1;

-- name: LookUpFeedByID :one

SELECT *
FROM FEEDS
WHERE id = $1;