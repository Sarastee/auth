-- +goose Up
-- +goose StatementBegin
create table if not exists users (
                                     id serial primary key,
                                     name text not null,
                                     email text unique not null,
                                     password text not null,
                                     role text not null,
                                     created_at timestamp with time zone default current_timestamp not null ,
                                     updated_at timestamp with time zone default current_timestamp not null
);

create index if not exists idx_role_email_name on users (role, email, name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
