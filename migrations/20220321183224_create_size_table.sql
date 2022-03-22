-- +goose Up
-- +goose StatementBegin
CREATE TABLE size(
                        id SERIAL,
                        name varchar(100) not null,
                        category_id bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE size;
-- +goose StatementEnd
