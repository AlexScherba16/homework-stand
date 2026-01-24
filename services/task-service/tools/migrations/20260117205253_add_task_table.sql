-- +goose Up
-- +goose StatementBegin
create table if not exists task(
    id      uuid primary key,
    user_id         bigint not null,
    category_id     text not null,
    comment         text,
    execution_time  interval not null,

    price_amount    bigint not null,

    status          varchar(32) not null,

    result          jsonb, -- новый JSON-результат

    created_at      timestamptz not null default now(),
    updated_at      timestamptz not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists task;
-- +goose StatementEnd
