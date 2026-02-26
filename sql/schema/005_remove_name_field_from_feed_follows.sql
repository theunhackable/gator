-- +goose Up

ALTER TABLE feed_follows
DROP COLUMN name;

-- +goose Down

ALTER TABLE feed_follows
ADD COLUMN name TEXT NOT NULL;


