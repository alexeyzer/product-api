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

type SizeQuery interface {
	Create(ctx context.Context, req datastruct.Size) (*datastruct.Size, error)
	Get(ctx context.Context, ID int64) (*datastruct.Size, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context) ([]*datastruct.Size, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type sizeQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *sizeQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.SizeTableName).
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

func (q *sizeQuery) List(ctx context.Context) ([]*datastruct.Size, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.SizeTableName).
		OrderBy("name")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var sizes []*datastruct.Size

	err = q.db.SelectContext(ctx, &sizes, query, args...)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (q *sizeQuery) Create(ctx context.Context, req datastruct.Size) (*datastruct.Size, error) {
	qb := q.builder.Insert(datastruct.SizeTableName).
		Columns(
			"name",
			"category",
		).
		Values(
			req.Name,
			req.Category,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var size datastruct.Size

	err = q.db.GetContext(ctx, &size, query, args...)
	if err != nil {
		return nil, err
	}

	return &size, nil
}

func (q *sizeQuery) Get(ctx context.Context, ID int64) (*datastruct.Size, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.SizeTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var size datastruct.Size

	err = q.db.GetContext(ctx, &size, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "size with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &size, nil
}

func (q *sizeQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.SizeTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var size datastruct.Size

	err = q.db.GetContext(ctx, &size, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewSizeQuery(db *sqlx.DB) SizeQuery {
	return &sizeQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
