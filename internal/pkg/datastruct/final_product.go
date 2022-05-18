package datastruct

const FinalProductTableName = "final_product"

type FinalProduct struct {
	ID        int64 `db:"id"`
	ProductID int64 `db:"product_id"`
	SizeID    int64 `db:"size_id"`
	Amount    int64 `db:"amount"`
	Sku       int64 `db:"sku"`
}

type FinalProductWithSizeName struct {
	ID        int64  `db:"id"`
	ProductID int64  `db:"product_id"`
	SizeID    int64  `db:"size_id"`
	SizeName  string `db:"size_name"`
	Amount    int64  `db:"amount"`
	Sku       int64  `db:"sku"`
}

type FullFinalProduct struct {
	ID           int64   `db:"id"`
	Amount       int64   `db:"amount"`
	Sku          int64   `db:"sku"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Url          string  `db:"url"`
	BrandName    string  `db:"brand_name"`
	CategoryName string  `db:"category_name"`
	Price        float64 `db:"price"`
	Color        string  `db:"color"`
	Size         string  `db:"size_name"`
}
