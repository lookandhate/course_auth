-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users
(
    id            serial primary key,
    name          text      not null,
    email         text      not null,
    password_hash text,
    role          smallint  not null,
    created_at    timestamp not null default now(),
    updated_at    timestamp
)

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table users;