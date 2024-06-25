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
	"strconv"
)

type NoteGetter interface {
	GetAllNotes(userID int64) ([]*notes.Note, error)
}

func NewNoteGetter(logger *log.Logger, noteGetter NoteGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.note.getnotes.NewNoteGetter"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		if err != nil {
			log.Error("failed get user id: ", "err", err)
		}

		noteGet := request.GetNotes{
			UserID: id,
		}

		log.Info("decoded query parameters", noteGet)

		if err := validator.New().Struct(noteGet); err != nil {
			validateErr := err.(validator.ValidationErrors)

			render.Status(r, http.StatusBadRequest)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		allNotes, err := noteGetter.GetAllNotes(noteGet.UserID)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get notes"))

			return
		}

		log.Info("successfully got notes")

		response.ResponseOKWithData(w, r, allNotes)
	}
}
