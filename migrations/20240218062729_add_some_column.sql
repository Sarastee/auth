-- +goose Up
-- +goose StatementBegin
create table if not exists users (
  id serial primary key,
  name text not null,
  email text not null,
  password text not null,
  password_confirm text not null,
  role text not null,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
