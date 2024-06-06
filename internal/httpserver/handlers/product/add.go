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

type ProductSaver interface {
	AddProduct(product products.Product) error
}

func NewProductSaver(logger *log.Logger, productSaver ProductSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.product.add.NewProductSaver"

		logger = log.With(
			"op: "+op,
			"request_id "+middleware.GetReqID(r.Context()),
		)

		var productAdd request.AddProduct
		var req request.Request = &productAdd
		request.Decode(w, r, &req)

		log.Info("decoded request body", productAdd)

		if err := validator.New().Struct(productAdd); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		product := products.Product{
			ProductID: 0,
			Name:      productAdd.Name,
			Carbs:     productAdd.Carbs,
		}

		err := productSaver.AddProduct(product)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to add product"))

			return
		}

		log.Info("added product")

		response.ResponseOK(w, r)
	}
}
