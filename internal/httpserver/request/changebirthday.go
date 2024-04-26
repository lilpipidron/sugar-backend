package request

import "time"

type ChangeBirthday struct {
	ID          int64     `json:"id"`
	NewBirthday time.Time `json:"new_birthday"`
}
