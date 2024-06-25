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

type BreadUnitsGetter interface {
	GetBreadUnitAmount(name string) (float64, error)
}

func NewBreadUnitsGetter(logger *log.Logger, carbsAmountGetter BreadUnitsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.product.getbreadunitsamount.NewBreadUnitsGetter"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		getBreadUnits := &request.GetBreadsUnit{
			Name: r.URL.Query().Get("name"),
		}

		log.Info("decoded query parameters", getBreadUnits)

		if err := validator.New().Struct(getBreadUnits); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		breadUnits, err := carbsAmountGetter.GetBreadUnitAmount(getBreadUnits.Name)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get bread units"))

			return
		}

		log.Info("successfully get bread units")
		response.ResponseOKWithData(w, r, breadUnits)
	}
}
