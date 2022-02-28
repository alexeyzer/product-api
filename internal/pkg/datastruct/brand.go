package datastruct

import "database/sql"

const BrandTableName = "brand"

type Brand struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Url         sql.NullString `db:"url"`
}
