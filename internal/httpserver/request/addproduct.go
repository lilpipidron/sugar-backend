package request

type AddProduct struct {
	ProductID  int64  `json:"id" default:"0"`
	Name       string `json:"name"`
	BreadUnits int    `json:"bread-units"`
}
