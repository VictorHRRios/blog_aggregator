-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
	INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)
	RETURNING *
)
SELECT
	inserted_feed_follow.*,
	feeds.name as feed_name,
	users.name as user_name
FROM inserted_feed_follow
JOIN users ON users.id = inserted_feed_follow.user_id
JOIN feeds ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT users.name as user_name, feeds.name as feed_name
FROM feed_follows
JOIN users on users.id = feed_follows.user_id
JOIN feeds on feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING users, feeds
WHERE users.id = feed_follows.user_id
AND feeds.id = feed_follows.feed_id
AND users.name = $1
AND feeds.url = $2;
