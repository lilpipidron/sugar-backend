package product

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

type CarbsAmountGetter interface {
	GetCarbsAmount(name string) (int, error)
}

func NewCarbsAmountGetter(logger *log.Logger, carbsAmountGetter CarbsAmountGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.product.getcarbsamount.NewCarbsAmountGetter"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		getCarbsAmount := &request.GetCarbsAmount{
			Name: r.URL.Query().Get("name"),
		}

		log.Info("decoded query parameters", getCarbsAmount)

		if err := validator.New().Struct(getCarbsAmount); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		carbs, err := carbsAmountGetter.GetCarbsAmount(getCarbsAmount.Name)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get carbs amount"))

			return
		}

		log.Info("successfully get carbs amount")
		response.ResponseOKWithData(w, r, carbs)
	}
}
