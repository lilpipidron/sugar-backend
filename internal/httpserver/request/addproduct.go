package request

type AddProduct struct {
	ProductID int    `json:"id"`
	Name      string `json:"name"`
	Carbs     int    `json:"carbs"`
}
