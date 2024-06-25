package products

type Product struct {
	ProductID  int64  `json:"id"`
	Name       string `json:"name"`
	BreadUnits int    `json:"bread-units"`
}
