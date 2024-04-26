package request

type ChangeCarbohydrateRatio struct {
	ID                   int64 `json:"id"`
	NewCarbohydrateRatio int   `json:"new_carbohydrate_ratio"`
}
