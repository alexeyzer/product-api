-- +goose Up
-- +goose StatementBegin
CREATE TYPE Content_type as ENUM ('PHOTO', 'VIDEO');

CREATE TABLE media(
                              id SERIAL,
                              product_id bigint not null,
                              url text,
                              content_type Content_type
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE media;

DROP TYPE Content_type;
-- +goose StatementEnd
