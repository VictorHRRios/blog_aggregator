-- +goose Up

alter table feeds
add last_fetched_at timestamp null;

-- +goose Down
alter table feeds
drop last_fetched_at;
