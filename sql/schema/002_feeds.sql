-- +goose Up

CREATE TABLE feeds(
	id UUID,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id UUID NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
