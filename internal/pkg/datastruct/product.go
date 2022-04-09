package datastruct

import "database/sql"

const ProductTableName = "product"

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Url         string  `db:"url"`
	BrandID     int64   `db:"brand_id"`
	CategoryID  int64   `db:"category_id"`
	Price       float64 `db:"price"`
	Color       string  `db:"color"`
}

type CreateProduct struct {
	Name        string
	Description string
	Image       []byte
	ContentType string
	BrandID     int64
	CategoryID  int64
	Price       float64
	Color       string
}

type ListProductRequest struct {
	Offset     int64
	Limit      int64
	CategoryID sql.NullInt64
	BrandID    sql.NullInt64
	Name       string
}
