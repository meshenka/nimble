-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS heroes (
    id BLOB PRIMARY KEY, -- UUID v7
    ancestry_name TEXT NOT NULL,
    class_name TEXT NOT NULL,
    motivation TEXT NOT NULL,
    origin TEXT NOT NULL,
    background_name TEXT NOT NULL,
    quirks TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS heroes;
-- +goose StatementEnd
