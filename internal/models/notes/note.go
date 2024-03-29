package notes

import "time"

type Note struct {
	NoteID     int64     `json:"id"`
	NoteType   NoteType  `json:"note-type"`
	DateTime   time.Time `json:"date-time"`
	SugarLevel int       `json:"sugar-level"`
	Products   []*NoteProduct
}
