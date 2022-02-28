-- +goose Up
-- +goose StatementBegin
CREATE TABLE brand(
                         id SERIAL,
                         name varchar(100) not null,
                         description TEXT,
                         url TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE brand;
-- +goose StatementEnd
