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

type ProductQuery interface {
	Create(ctx context.Context, req datastruct.Product) (*datastruct.Product, error)
	Get(ctx context.Context, ID int64) (*datastruct.Product, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context, req datastruct.ListProductRequest) ([]*datastruct.Product, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type productQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *productQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.ProductTableName).
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

func (q *productQuery) List(ctx context.Context, req datastruct.ListProductRequest) ([]*datastruct.Product, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ProductTableName).
		Offset(uint64(req.Offset)).
		Limit(uint64(req.Limit))

	if req.Name != "" {
		qb = qb.Where(squirrel.ILike{"name": "%" + req.Name + "%"})
	}
	if req.BrandID.Valid {
		qb = qb.Where(squirrel.Eq{"brand_id": req.BrandID.Int64})
	}
	if req.CategoryID.Valid {
		qb = qb.Where(squirrel.Eq{"category_id": req.CategoryID.Int64})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var products []*datastruct.Product

	err = q.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (q *productQuery) Create(ctx context.Context, req datastruct.Product) (*datastruct.Product, error) {
	qb := q.builder.Insert(datastruct.ProductTableName).
		Columns(
			"name",
			"description",
			"url",
			"brand_id",
			"category_id",
			"price",
			"color",
		).
		Values(
			req.Name,
			req.Description,
			req.Url,
			req.BrandID,
			req.CategoryID,
			req.Price,
			req.Color,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var product datastruct.Product

	err = q.db.GetContext(ctx, &product, query, args...)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (q *productQuery) Get(ctx context.Context, ID int64) (*datastruct.Product, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ProductTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var product datastruct.Product

	err = q.db.GetContext(ctx, &product, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "product with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &product, nil
}

func (q *productQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ProductTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var product datastruct.Product

	err = q.db.GetContext(ctx, &product, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewProductQuery(db *sqlx.DB) ProductQuery {
	return &productQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
