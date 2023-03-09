-- +goose Up
-- +goose StatementBegin
create table account (
    id bigserial primary key ,
    name text not null,
    status text not null default 'new',
    access_level text not null default 'none',
    opened_at timestamptz not null default now(),
    closed_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
