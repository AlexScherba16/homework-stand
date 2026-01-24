-- +goose Up
-- +goose StatementBegin
create table if not exists profile(
    user_id     bigint primary key,
    name text not null,
    email text not null,
    is_task_allowed boolean default true,
    created_at timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists profile;
-- +goose StatementEnd
