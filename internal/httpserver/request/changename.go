package request

type ChangeName struct {
	ID      int64  `json:"id"`
	NewName string `json:"new_name"`
}
