-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_in_order(
                      id SERIAL,
                      final_product_id bigint not null,
                      amount bigint not null,
                      order_id bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_in_order;
-- +goose StatementEnd
