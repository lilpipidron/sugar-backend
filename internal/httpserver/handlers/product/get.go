package product

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/products"
	"net/http"
)

type ProductsGetter interface {
	GetProductsWithValueInName(value string) ([]*products.Product, error)
}

func NewProductsGetter(logger *log.Logger, productsGetter ProductsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.product.get.NewProductsGetter"

		logger = log.With(
			"op: "+op,
			"request_id "+middleware.GetReqID(r.Context()),
		)

		productsGet := request.GetProducts{
			Name: r.URL.Query().Get("name"),
		}

		log.Info("decoded query parameters", productsGet)

		if err := validator.New().Struct(productsGet); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		productsArr, err := productsGetter.GetProductsWithValueInName(productsGet.Name)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get product"))

			return
		}

		log.Info("successfully get products")

		response.ResponseOKWithData(w, r, productsArr)
	}
}
