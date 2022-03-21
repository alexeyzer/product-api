package repository

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MediaQuery interface {
	Create(ctx context.Context, req datastruct.Media) (*datastruct.Media, error)
	Get(ctx context.Context, ID int64) (*datastruct.Media, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, productID int64) ([]*datastruct.Media, error)
}

type mediaQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *mediaQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.MediaTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = q.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (q *mediaQuery) List(ctx context.Context, productID int64) ([]*datastruct.Media, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.MediaTableName).
		Where(squirrel.Eq{"product_id": productID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var medias []*datastruct.Media

	err = q.db.SelectContext(ctx, &medias, query, args...)
	if err != nil {
		return nil, err
	}

	return medias, nil
}

func (q *mediaQuery) Create(ctx context.Context, req datastruct.Media) (*datastruct.Media, error) {
	qb := q.builder.Insert(datastruct.MediaTableName).
		Columns(
			"url",
			"product_id",
			"content_type",
		).
		Values(
			req.Url,
			req.ProductID,
			req.ContentType,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var media datastruct.Media

	err = q.db.GetContext(ctx, &media, query, args...)
	if err != nil {
		return nil, err
	}

	return &media, nil
}

func (q *mediaQuery) Get(ctx context.Context, ID int64) (*datastruct.Media, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.MediaTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var media datastruct.Media

	err = q.db.GetContext(ctx, &media, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "media with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &media, nil
}

func NewMediaQuery(db *sqlx.DB) MediaQuery {
	return &mediaQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
