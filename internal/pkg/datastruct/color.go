package datastruct

const ColorTableName = "color"

type Color struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
