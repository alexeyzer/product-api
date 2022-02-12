package datastruct

import "database/sql"

const CategoryTableName = "category"

type Category struct {
	ID       int64         `db:"id"`
	Name     string        `db:"name"`
	Level    int64         `db:"level"`
	ParentID sql.NullInt64 `db:"parent_id"`
}

type ListCategoryRequest struct {
	Offset int64
	Limit  int64
	Level  sql.NullInt64
	Name   sql.NullString
}
