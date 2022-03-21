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

type BrandQuery interface {
	Create(ctx context.Context, req datastruct.Brand) (*datastruct.Brand, error)
	Get(ctx context.Context, ID int64) (*datastruct.Brand, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context) ([]*datastruct.Brand, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type brandQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *brandQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.BrandTableName).
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

func (q *brandQuery) List(ctx context.Context) ([]*datastruct.Brand, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.BrandTableName).
		OrderBy("name")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var brands []*datastruct.Brand

	err = q.db.SelectContext(ctx, &brands, query, args...)
	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (q *brandQuery) Create(ctx context.Context, req datastruct.Brand) (*datastruct.Brand, error) {
	qb := q.builder.Insert(datastruct.BrandTableName).
		Columns(
			"name",
			"description",
			"url",
		).
		Values(
			req.Name,
			req.Description,
			req.Url,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var brand datastruct.Brand

	err = q.db.GetContext(ctx, &brand, query, args...)
	if err != nil {
		return nil, err
	}

	return &brand, nil
}

func (q *brandQuery) Get(ctx context.Context, ID int64) (*datastruct.Brand, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.BrandTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var brand datastruct.Brand

	err = q.db.GetContext(ctx, &brand, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "brand with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &brand, nil
}

func (q *brandQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.BrandTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var brand datastruct.Brand

	err = q.db.GetContext(ctx, &brand, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewBrandQuery(db *sqlx.DB) BrandQuery {
	return &brandQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
