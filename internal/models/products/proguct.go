package products

type Product struct {
	ProductID int64  `json:"id"`
	Name      string `json:"name"`
	Carbs     int    `json:"carbs"`
}
