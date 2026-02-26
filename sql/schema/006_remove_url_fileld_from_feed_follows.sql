-- +goose Up

ALTER TABLE feed_follows
DROP COLUMN url;

-- +goose Down

ALTER TABLE feed_follows
ADD COLUMN url TEXT NOT NULL;

