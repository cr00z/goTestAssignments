-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS playlist
(
    id bigint NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    duration bigint NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS playlist;
-- +goose StatementEnd
