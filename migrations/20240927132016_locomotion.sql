-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id bigserial primary key,
    first_name varchar(255) not null,
    second_name varchar(255) not null,
    patronymic varchar(255) default null,
    age int not null,
    sex varchar(10) not null,
    date_birth  timestamp default now() not null,
    created_at timestamp default now() not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
