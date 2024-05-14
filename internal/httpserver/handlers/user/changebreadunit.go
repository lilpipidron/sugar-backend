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

type BreadUnitChanger interface {
	ChangeBreadUnit(id int64, newBreadUnit int) error
}

func NewBreadUnitChanger(logger *log.Logger, breadUnitChanger BreadUnitChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changebreadunit.NewBreadUnitChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeBreadUnit request.ChangeBreadUnit
		var req request.Request = &changeBreadUnit
		request.Decode(w, r, &req)

		log.Info("decoded request body", changeBreadUnit)

		if err := validator.New().Struct(changeBreadUnit); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := breadUnitChanger.ChangeBreadUnit(changeBreadUnit.ID, changeBreadUnit.NewBreadUnit)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change bread unit"))

			return
		}

		log.Info("changed bread unit")

		response.ResponseOK(w, r)
	}
}
