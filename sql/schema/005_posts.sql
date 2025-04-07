-- +goose Up
create table posts(
	id uuid,
	created_at timestamp not null,
	updated_at timestamp not null,
	title text not null,
	url text unique not null,
	description text not null,
	published_at text not null,
	feed_id uuid not null,
	primary key (id),
	foreign key (feed_id) references feeds(id) on delete cascade
);

-- +goose Down
drop table posts;
