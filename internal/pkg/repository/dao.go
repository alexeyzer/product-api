package repository

import (
	"fmt"
	"github.com/alexeyzer/product-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DAO interface {
	CategoryQuery() CategoryQuery
	BrandQuery() BrandQuery
	ColorQuery() ColorQuery
	SizeQuery() SizeQuery
	MediaQuery() MediaQuery
}

type dao struct {
	brandQuery    BrandQuery
	categoryQuery CategoryQuery
	colorQuery    ColorQuery
	sizeQuery     SizeQuery
	mediaQuery    MediaQuery
	db            *sqlx.DB
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

func (d *dao) BrandQuery() BrandQuery {
	if d.brandQuery == nil {
		d.brandQuery = NewBrandQuery(d.db)
	}
	return d.brandQuery
}

func (d *dao) ColorQuery() ColorQuery {
	if d.colorQuery == nil {
		d.colorQuery = NewColorQuery(d.db)
	}
	return d.colorQuery
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
