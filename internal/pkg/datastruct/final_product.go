package datastruct

const FinalProductTableName = "final_product"

type FinalProduct struct {
	ID        int64 `db:"id"`
	ProductID int64 `db:"product_id"`
	SizeID    int64 `db:"size_id"`
	ColorID   int64 `db:"color_id"`
	Amount    int64 `db:"amount"`
	Sku       int64 `db:"sku"`
	Price     int64 `db:"price"`
}
