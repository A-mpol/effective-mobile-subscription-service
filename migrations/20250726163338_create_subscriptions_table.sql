-- +goose Up
-- +goose StatementBegin
create table subscriptions (
    id uuid not null primary key,
    user_id uuid not null,
    service_name text not null,
    price serial,
    start_date timestamp not null default now(),
    end_date timestamp,
    deleted_at timestamp 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table subscriptions;
-- +goose StatementEnd
