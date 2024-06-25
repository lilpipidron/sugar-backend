package note

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"net/http"
)

type NoteDeleter interface {
	DeleteNote(noteID int64) error
}

func NewNoteDelete(logger *log.Logger, noteDeleter NoteDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.note.delete.NewNoteDelete"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var noteDelete request.DeleteNote
		var req request.Request = &noteDelete
		request.Decode(w, r, &req)

		log.Info("decoded request body", noteDelete)

		if err := validator.New().Struct(noteDelete); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := noteDeleter.DeleteNote(noteDelete.NoteID)
		if err != nil {
			log.Error(err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error("failed to delete note"))

			return
		}

		log.Info("delete note")

		response.ResponseOK(w, r)
	}
}
