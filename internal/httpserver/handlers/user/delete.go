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

type UserDeleter interface {
	DeleteUser(id int64) error
}

func NewUserDelete(logger *log.Logger, userDeleter UserDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.delete.NewUserDelete"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var userDelete request.DeleteUser
		var req request.Request = &userDelete
		request.Decode(w, r, &req)
		userDelete = (req).(request.DeleteUser)

		log.Info("decoded request body", userDelete)

		if err := validator.New().Struct(userDelete); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := userDeleter.DeleteUser(userDelete.Id)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to delete user"))

			return
		}

		log.Info("delete user")

		response.ResponseOK(w, r)
	}
}
