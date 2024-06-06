package request

import (
	"github.com/lilpipidron/sugar-backend/internal/models/notes"
	"time"
)

type AddNote struct {
	UserID     int64          `json:"user-id"`
	NoteID     int64          `json:"note-id"`
	NoteType   notes.NoteType `json:"note-type"`
	DateTime   time.Time      `json:"date-time"`
	SugarLevel int            `json:"sugar-level"`
	Products   []*notes.NoteProduct
}
