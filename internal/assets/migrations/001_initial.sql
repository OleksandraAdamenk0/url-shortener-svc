-- +migrate Up
CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    original_url TEXT NOT NULL,
    short_url TEXT UNIQUE NOT NULL
);
-- +migrate Down
DROP TABLE urls;
