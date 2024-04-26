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

type WeightChanger interface {
	ChangeWeight(id int64, newWeight int) error
}

func NewWeightChanger(logger *log.Logger, weightChanger WeightChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changeweight.NewWeightChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeWeight request.ChangeWeight
		var req request.Request = &changeWeight
		request.Decode(w, r, &req)
		changeWeight = (req).(request.ChangeWeight)

		log.Info("decoded request body", changeWeight)

		if err := validator.New().Struct(changeWeight); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := weightChanger.ChangeWeight(changeWeight.ID, changeWeight.NewWeight)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change weight"))

			return
		}

		log.Info("changed weight")

		response.ResponseOK(w, r)
	}
}
