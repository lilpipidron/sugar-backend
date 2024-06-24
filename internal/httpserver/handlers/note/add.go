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
)

type NoteSaver interface {
	AddNote(note notes.Note, userID int64) error
}

func NewNoteSaver(logger *log.Logger, noteSaver NoteSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.note.add.NewNoteAdd"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var noteAdd request.AddNote
		var req request.Request = &noteAdd
		request.Decode(w, r, &req)

		log.Info("decoded request body", noteAdd)

		if err := validator.New().Struct(noteAdd); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		note := notes.Note{
			NoteID:     noteAdd.NoteID,
			DateTime:   noteAdd.DateTime,
			SugarLevel: noteAdd.SugarLevel,
		}

		err := noteSaver.AddNote(note, noteAdd.UserID)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to add note"))

			return
		}

		log.Info("added note")

		response.ResponseOK(w, r)
	}
}
