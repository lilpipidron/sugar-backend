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

		date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
		if err != nil {
			log.Error("failed to parse date", "date", r.URL.Query().Get("date"))
		}
		noteGetByDate := request.GetNotesByDate{
			DateTime: date,
		}

		log.Info("decoded query parameters", noteGetByDate)

		if err := validator.New().Struct(noteGetByDate); err != nil {
			validateErr := err.(validator.ValidationErrors)

			render.Status(r, http.StatusBadRequest)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		allNotes, err := noteGetter.GetNotesByDate(noteGetByDate.UserID, noteGetByDate.DateTime)
		if err != nil {
			log.Error(err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error("failed to get notes by date"))

			return
		}

		log.Info("successfully got notes by date")

		response.ResponseOKWithData(w, r, allNotes)
	}
}
