-- +goose Up
-- +goose StatementBegin
CREATE TABLE final_product(
                      id SERIAL,
                      sku bigint not null,
                      product_id bigint not null,
                      size_id bigint not null,
                      amount bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE final_product;
-- +goose StatementEnd
