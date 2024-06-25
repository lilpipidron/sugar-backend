package products

type Product struct {
	ProductID  int64   `json:"id"`
	Name       string  `json:"name"`
	BreadUnits float64 `json:"bread-units"`
}
