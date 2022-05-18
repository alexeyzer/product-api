-- +goose Up
-- +goose StatementBegin
CREATE TABLE product(
                      id SERIAL,
                      name varchar(100) not null,
                      description TEXT,
                      url TEXT,
                      brand_id bigint not null,
                      category_id bigint not null,
                      price double precision not null,
                      color TEXT not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
