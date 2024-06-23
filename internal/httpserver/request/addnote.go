package request

import (
	"encoding/json"
	"github.com/lilpipidron/sugar-backend/internal/models/notes"
	"time"
)

type AddNote struct {
	UserID     int64          `json:"user-id"`
	NoteID     int64          `json:"note-id" default:"0"`
	NoteType   notes.NoteType `json:"note-type"`
	DateTime   time.Time      `json:"date-time"`
	SugarLevel int            `json:"sugar-level"`
}

func (u *AddNote) UnmarshalJSON(data []byte) error {
	type Alias AddNote
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
