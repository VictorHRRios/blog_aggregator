-- name: CreatePost :one
insert into posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
values (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8
)
on conflict (url) do nothing
returning *;

-- name: GetPostForUser :many
select posts.*
from posts
join feeds on feeds.id = posts.feed_id
join users on users.id = feeds.user_id
where users.id = $1
order by posts.published_at desc
limit $2;
