-- +goose Up
-- +goose StatementBegin
insert into account(id, name, status, access_level, opened_at)
values (1, 'test_account','open','full',now());
ALTER SEQUENCE account_id_seq RESTART WITH 10;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
