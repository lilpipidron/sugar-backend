package product

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/products"
	"net/http"
)

type AllProductsGetter interface {
	GetAllProducts() ([]*products.Product, error)
}

func NewAllProductsGetter(logger *log.Logger, productsGetter AllProductsGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.product.get.NewAllProductsGetter"

		logger = log.With(
			"op: "+op,
			"request_id "+middleware.GetReqID(r.Context()),
		)

		productsArr, err := productsGetter.GetAllProducts()
		if err != nil {
			log.Error(err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error("failed to get product"))

			return
		}

		log.Info("successfully get products")

		response.ResponseOKWithData(w, r, productsArr)
	}
}
