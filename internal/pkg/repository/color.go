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

type ColorQuery interface {
	Create(ctx context.Context, req datastruct.Color) (*datastruct.Color, error)
	Get(ctx context.Context, ID int64) (*datastruct.Color, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context) ([]*datastruct.Color, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type colorQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *colorQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.ColorTableName).
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

func (q *colorQuery) List(ctx context.Context) ([]*datastruct.Color, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ColorTableName).
		OrderBy("name")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var colors []*datastruct.Color

	err = q.db.SelectContext(ctx, &colors, query, args...)
	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (q *colorQuery) Create(ctx context.Context, req datastruct.Color) (*datastruct.Color, error) {
	qb := q.builder.Insert(datastruct.ColorTableName).
		Columns(
			"name",
		).
		Values(
			req.Name,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var color datastruct.Color

	err = q.db.GetContext(ctx, &color, query, args...)
	if err != nil {
		return nil, err
	}

	return &color, nil
}

func (q *colorQuery) Get(ctx context.Context, ID int64) (*datastruct.Color, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ColorTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var color datastruct.Color

	err = q.db.GetContext(ctx, &color, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "color with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &color, nil
}

func (q *colorQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ColorTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var color datastruct.Color

	err = q.db.GetContext(ctx, &color, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewColorQuery(db *sqlx.DB) ColorQuery {
	return &colorQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
