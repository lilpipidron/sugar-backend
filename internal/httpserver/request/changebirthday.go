package request

import (
	"encoding/json"
	"time"
)

type ChangeBirthday struct {
	ID          int64     `json:"id"`
	NewBirthday time.Time `json:"new_birthday"`
}

func (u *ChangeBirthday) UnmarshalJSON(data []byte) error {
	type Alias ChangeBirthday
	aux := &struct {
		NewBirthday string `json:"new_birthday"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	u.NewBirthday, err = time.Parse("2006-01-02", aux.NewBirthday)
	return err
}
