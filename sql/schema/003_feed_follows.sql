-- +goose Up
CREATE TABLE feed_follows(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL,
  feed_id UUID NOT NULL,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL ,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT fk_feed
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE feed_follows;
