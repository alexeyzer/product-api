-- +goose Up
-- +goose StatementBegin
CREATE TABLE product(
                      id SERIAL,
                      name varchar(100) not null,
                      description TEXT,
                      url TEXT,
                      brand_id bigint,
                      category_id bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
