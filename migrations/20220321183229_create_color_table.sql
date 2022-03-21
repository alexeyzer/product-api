-- +goose Up
-- +goose StatementBegin
CREATE TABLE color(
                     id SERIAL,
                     name varchar(100) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE color;
-- +goose StatementEnd
