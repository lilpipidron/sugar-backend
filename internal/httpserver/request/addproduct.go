package request

type AddProduct struct {
	ProductID int64  `json:"id"`
	Name      string `json:"name"`
	Carbs     int    `json:"carbs"`
}
