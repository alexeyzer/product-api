package datastruct

type FullProduct struct {
	ID          int64
	Name        string
	Description string
	Url         string
	BrandID     int64
	CategoryID  int64
	Price       float64
	Color       string
	Sizes       []*Size
}
