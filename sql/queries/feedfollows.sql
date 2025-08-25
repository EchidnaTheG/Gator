-- name: CreateFeedFollow :one


WITH inserted_feed_follow AS(
INSERT INTO feed_follows ( id, created_at, updated_at, userid, feedid)
VALUES($1,$2,$3,$4,$5)
RETURNING *
)
SELECT 
    inserted_feed_follow.*, feeds.name AS feed_name, users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.userid=users.id
INNER JOIN feeds
ON inserted_feed_follow.feedid = feeds.id;

-- name: GetFeedFollowsForUser :many

SELECT *
FROM feed_follows
WHERE userid = $1;
