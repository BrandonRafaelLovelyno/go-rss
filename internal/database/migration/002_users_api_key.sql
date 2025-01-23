-- +goose Up

ALTER TABLE users ADD COLUMN api_key VARCHAR(255) NOT NULL DEFAULT (
    encode(random()::text::bytea, 'hex')
);

-- +goose Down

DROP TABLE users;

