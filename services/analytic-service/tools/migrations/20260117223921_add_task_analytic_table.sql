-- +goose Up
-- +goose StatementBegin
create table if not exists task_analytic (
    task_id        uuid primary key,
    user_id        bigint not null,
    category_id    text not null,
    status         text not null,
    created_at     timestamptz not null,
    completed_at   timestamptz,
    comment        text,
    execution_time interval not null,
    price_amount   numeric not null,
    updated_at     timestamptz not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists task_analytic;
-- +goose StatementEnd
