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
	GetByProductID(ctx context.Context, productID int64) ([]*datastruct.Size, error)
	Delete(ctx context.Context, ID int64) error
	List(ctx context.Context) ([]*datastruct.SizeWithCategoryName, error)
	Exists(ctx context.Context, name string) (bool, error)
	Update(ctx context.Context, req datastruct.Size) (*datastruct.Size, error)
}

type sizeQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *sizeQuery) Update(ctx context.Context, req datastruct.Size) (*datastruct.Size, error) {
	qb := q.builder.Update(datastruct.SizeTableName).
		Set("name", req.Name).
		Set("category_id", req.CategoryID).
		Where(squirrel.Eq{"id": req.ID}).
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

func (q *sizeQuery) GetByProductID(ctx context.Context, productID int64) ([]*datastruct.Size, error) {
	qb := q.builder.
		Select(
			"st.id",
			"st.name",
			"st.category_id",
		).
		From(datastruct.SizeTableName + " as st").
		LeftJoin(datastruct.FinalProductTableName + " as fpt on fpt.size_id = st.id").
		Where(squirrel.Eq{"fpt.id": productID})
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

func (q *sizeQuery) List(ctx context.Context) ([]*datastruct.SizeWithCategoryName, error) {
	qb := q.builder.
		Select("stn.id", "stn.name", "ctn.name as category_name").
		From(datastruct.SizeTableName + " as stn").
		OrderBy("category_id").
		LeftJoin(datastruct.CategoryTableName + " as ctn on ctn.id = stn.category_id")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var sizes []*datastruct.SizeWithCategoryName

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
			"category_id",
		).
		Values(
			req.Name,
			req.CategoryID,
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
