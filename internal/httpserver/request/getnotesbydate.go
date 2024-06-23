package request

import (
	"encoding/json"
	"time"
)

type GetNotesByDate struct {
	UserID   int64     `json:"id"`
	DateTime time.Time `json:"date-time"`
}

func (u *GetNotesByDate) UnmarshalJSON(data []byte) error {
	type Alias GetNotesByDate
	aux := &struct {
		Birthday string `json:"date-time"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	u.DateTime, err = time.Parse("2006-01-02", aux.Birthday)
	return err
}
