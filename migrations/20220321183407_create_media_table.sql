-- +goose Up
-- +goose StatementBegin
CREATE TABLE media(
                              id SERIAL,
                              product_id bigint not null,
                              url text,
                              content_type text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE media;
-- +goose StatementEnd
