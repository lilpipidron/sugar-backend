package user

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
)

type NameChanger interface {
	ChangeName(id int64, newName string) error
}

func NewNameChanger(logger *log.Logger, nameChanger NameChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changeName.NewNameChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeName request.ChangeName
		var req request.Request = &changeName
		request.Decode(w, r, &req)

		log.Info("decoded request body", changeName)

		if err := validator.New().Struct(changeName); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := nameChanger.ChangeName(changeName.ID, changeName.NewName)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change name"))

			return
		}

		log.Info("changed name")

		response.ResponseOK(w, r)
	}
}
