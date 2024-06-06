package request

import "time"

type GetNotesByDate struct {
	UserID   int64     `json:"id"`
	DateTime time.Time `json:"date-time"`
}
