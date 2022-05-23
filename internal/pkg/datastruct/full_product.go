package datastruct

type FullProduct struct {
	ID           int64   `db:"id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Url          string  `db:"url"`
	BrandName    string  `db:"brand_name"`
	CategoryName string  `db:"category_name"`
	Price        float64 `db:"price"`
	Color        string  `db:"color"`
	Sizes        []*Size
	IsFavorite   bool
	UserQuantity int64
	FavoriteID   int64
}
