package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type CategoryQuery interface {
	Create(ctx context.Context, req datastruct.Category) (*datastruct.Category, error)
	Update(ctx context.Context, req datastruct.Category) (*datastruct.Category, error)
	List(ctx context.Context, req datastruct.ListCategoryRequest) ([]datastruct.Category, error)
	Delete(ctx context.Context, ID int64) error
	Get(ctx context.Context, ID int64) (*datastruct.Category, error)
	Exists(ctx context.Context, name string) (bool, error)
}

type categoryQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *categoryQuery) List(ctx context.Context, req datastruct.ListCategoryRequest) ([]datastruct.Category, error) {
	qb := q.builder.
		Select("*").
		Offset(uint64(req.Offset)).
		Limit(uint64(req.Limit)).
		From(datastruct.CategoryTableName)

	if req.Name.Valid {
		qb = qb.Where(squirrel.ILike{"name": "%" + strings.ToLower(req.Name.String) + "%"})
	}
	if req.Level.Valid {
		qb = qb.Where(squirrel.Eq{"level": req.Level.Int64})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var categories []datastruct.Category

	err = q.db.SelectContext(ctx, &categories, query, args...)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (q *categoryQuery) Delete(ctx context.Context, ID int64) error {
	qb := q.builder.
		Delete(datastruct.CategoryTableName).
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

func (q *categoryQuery) Update(ctx context.Context, req datastruct.Category) (*datastruct.Category, error) {
	qb := q.builder.Update(datastruct.CategoryTableName).
		Set("name", req.Name).
		Set("level", req.Level).
		Set("parent_id", req.ParentID).
		Where(squirrel.Eq{"id": req.ID}).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var category datastruct.Category

	err = q.db.GetContext(ctx, &category, query, args...)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (q *categoryQuery) Create(ctx context.Context, req datastruct.Category) (*datastruct.Category, error) {
	fmt.Println(req)
	qb := q.builder.Insert(datastruct.CategoryTableName).
		Columns(
			"name",
			"level",
			"parent_id",
		).
		Values(
			req.Name,
			req.Level,
			req.ParentID,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var category datastruct.Category

	err = q.db.GetContext(ctx, &category, query, args...)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (q *categoryQuery) Get(ctx context.Context, ID int64) (*datastruct.Category, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.CategoryTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var category datastruct.Category

	err = q.db.GetContext(ctx, &category, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "category with ID = %d doesn't exist", ID)
		}
		return nil, err
	}

	return &category, nil
}

func (q *categoryQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.CategoryTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var category datastruct.Category

	err = q.db.GetContext(ctx, &category, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewCategoryQuery(db *sqlx.DB) CategoryQuery {
	return &categoryQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
