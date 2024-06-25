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

type CarbohydrateRatioChanger interface {
	ChangeCarbohydrateRatio(id int64, newCarbohydrateRatio int) error
}

func NewCarbohydrateRatioChanger(logger *log.Logger, carbohydrateRatioChanger CarbohydrateRatioChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changecarbohydrateratio.NewCarbohydrateRatioChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeCarbohydrateRatio request.ChangeCarbohydrateRatio
		var req request.Request = &changeCarbohydrateRatio
		request.Decode(w, r, &req)

		log.Info("decoded request body", changeCarbohydrateRatio)

		if err := validator.New().Struct(changeCarbohydrateRatio); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := carbohydrateRatioChanger.ChangeCarbohydrateRatio(changeCarbohydrateRatio.ID, changeCarbohydrateRatio.NewCarbohydrateRatio)
		if err != nil {
			log.Error(err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error("failed to change carbohydrate ratio"))

			return
		}

		log.Info("changed carbohydrate ratio")

		response.ResponseOK(w, r)
	}
}
