package note

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/notes"
	"net/http"
	"time"
)

type NoteGetterByDate interface {
	GetNotesByDate(userID int64, dateTime time.Time) ([]*notes.Note, error)
}

func NewNoteGetterByDate(logger *log.Logger, noteGetter NoteGetterByDate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.note.getnotesbydate.NewNoteGetter"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var noteGetByDate request.GetNotesByDate
		var req request.Request = &noteGetByDate
		request.Decode(w, r, &req)

		log.Info("decoded request body", noteGetByDate)

		if err := validator.New().Struct(noteGetByDate); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		allNotes, err := noteGetter.GetNotesByDate(noteGetByDate.UserID, noteGetByDate.DateTime)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get notes by date"))

			return
		}

		log.Info("successfully got notes by date")

		response.ResponseOKWithData(w, r, allNotes)
	}
}
