package notes

import "time"

type Note struct {
	UserID     int64
	NoteType   NoteType
	DateTime   time.Time
	SugarLevel int
	Products   []NoteProduct
}
