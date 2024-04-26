package request

type ChangeWeight struct {
	ID        int64 `json:"id"`
	NewWeight int   `json:"new_weight"`
}
