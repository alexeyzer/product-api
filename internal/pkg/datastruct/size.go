package datastruct

const SizeTableName = "size"

type Size struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	CategoryID int64  `db:"category_id"`
}

type SizeWithCategoryName struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	CategoryName string `db:"category_name"`
}
