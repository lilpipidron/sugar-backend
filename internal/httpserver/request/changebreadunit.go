package request

type ChangeBreadUnit struct {
	ID           int64 `json:"id"`
	NewBreadUnit int   `json:"new_bread_unit"`
}
