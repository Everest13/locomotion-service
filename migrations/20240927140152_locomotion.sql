-- +goose Up
-- +goose StatementBegin
alter table if exists users
add column if not exists birthday timestamp default null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users
drop column if exists birthday;
-- +goose StatementEnd
