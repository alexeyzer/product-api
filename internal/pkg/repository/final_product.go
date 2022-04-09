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

type FinalProductQuery interface {
	Create(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error)
	Get(ctx context.Context, ID int64) (*datastruct.FinalProduct, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, productID int64) ([]*datastruct.FinalProduct, error)
	Exists(ctx context.Context, sku int64) (bool, error)
	DeleteByProductID(ctx context.Context, productID int64) error
}

type finalProductQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *finalProductQuery) DeleteByProductID(ctx context.Context, productID int64) error {
	qb := q.builder.
		Delete(datastruct.FinalProductTableName).
		Where(squirrel.Eq{"product_id": productID})
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

func (q *finalProductQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.FinalProductTableName).
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

func (q *finalProductQuery) List(ctx context.Context, productID int64) ([]*datastruct.FinalProduct, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.FinalProductTableName).
		Where(squirrel.Eq{"product_id": productID})

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var finalProduct []*datastruct.FinalProduct

	err = q.db.SelectContext(ctx, &finalProduct, query, args...)
	if err != nil {
		return nil, err
	}

	return finalProduct, nil
}

func (q *finalProductQuery) Create(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error) {
	qb := q.builder.Insert(datastruct.FinalProductTableName).
		Columns(
			"product_id",
			"size_id",
			"amount",
			"sku",
		).
		Values(
			req.ProductID,
			req.SizeID,
			req.Amount,
			req.Sku,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var finalProduct datastruct.FinalProduct

	err = q.db.GetContext(ctx, &finalProduct, query, args...)
	if err != nil {
		return nil, err
	}

	return &finalProduct, nil
}

func (q *finalProductQuery) Get(ctx context.Context, ID int64) (*datastruct.FinalProduct, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.FinalProductTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var finalProduct datastruct.FinalProduct

	err = q.db.GetContext(ctx, &finalProduct, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "finalProduct with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &finalProduct, nil
}

func (q *finalProductQuery) Exists(ctx context.Context, sku int64) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.FinalProductTableName).
		Where(squirrel.Eq{"sku": sku})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var finalProduct datastruct.FinalProduct

	err = q.db.GetContext(ctx, &finalProduct, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewFinalProductQuery(db *sqlx.DB) FinalProductQuery {
	return &finalProductQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
