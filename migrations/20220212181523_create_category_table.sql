-- +goose Up
-- +goose StatementBegin
CREATE TABLE category(
                      id SERIAL,
                      name varchar(100) not null,
                      level bigint not null,
                      parent_id bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE category;
-- +goose StatementEnd
