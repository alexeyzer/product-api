-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders(
                      id SERIAL,
                      user_id bigint not null,
                      status text,
                      order_date time,
                      order_cost bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
