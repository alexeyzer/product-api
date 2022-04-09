package repository

import (
	"context"
	"fmt"
	"github.com/alexeyzer/product-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DAO interface {
	CategoryQuery() CategoryQuery
	BrandQuery() BrandQuery
	SizeQuery() SizeQuery
	MediaQuery() MediaQuery
	ProductQuery() ProductQuery
	FinalProductQuery() FinalProductQuery
	BeginTransaction(ctx context.Context) (*sqlx.Tx, error)
}

type dao struct {
	brandQuery        BrandQuery
	categoryQuery     CategoryQuery
	sizeQuery         SizeQuery
	mediaQuery        MediaQuery
	productQuery      ProductQuery
	finalProductQuery FinalProductQuery
	db                *sqlx.DB
}

func NewDao() (DAO, error) {
	dao := &dao{}
	dbConf := config.Config.Database
	dsn := fmt.Sprintf(dbConf.Dsn,
		dbConf.Host,
		dbConf.Port,
		dbConf.Dbname,
		dbConf.User,
		dbConf.Password,
		dbConf.Ssl)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	dao.db = conn
	return dao, nil
}

func (d *dao) BeginTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return d.db.BeginTxx(ctx, nil)
}

func (d *dao) BrandQuery() BrandQuery {
	if d.brandQuery == nil {
		d.brandQuery = NewBrandQuery(d.db)
	}
	return d.brandQuery
}

func (d *dao) CategoryQuery() CategoryQuery {
	if d.categoryQuery == nil {
		d.categoryQuery = NewCategoryQuery(d.db)
	}
	return d.categoryQuery
}

func (d *dao) SizeQuery() SizeQuery {
	if d.sizeQuery == nil {
		d.sizeQuery = NewSizeQuery(d.db)
	}
	return d.sizeQuery
}

func (d *dao) MediaQuery() MediaQuery {
	if d.mediaQuery == nil {
		d.mediaQuery = NewMediaQuery(d.db)
	}
	return d.mediaQuery
}

func (d *dao) ProductQuery() ProductQuery {
	if d.productQuery == nil {
		d.productQuery = NewProductQuery(d.db)
	}
	return d.productQuery
}

func (d *dao) FinalProductQuery() FinalProductQuery {
	if d.finalProductQuery == nil {
		d.finalProductQuery = NewFinalProductQuery(d.db)
	}
	return d.finalProductQuery
}
