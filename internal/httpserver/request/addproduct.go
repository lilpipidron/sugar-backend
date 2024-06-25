package request

type AddProduct struct {
	ProductID  int64   `json:"id" default:"0"`
	Name       string  `json:"name"`
	BreadUnits float64 `json:"bread-units"`
}
