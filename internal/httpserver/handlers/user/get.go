package user

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/users"
	"net/http"
)

type UserGetter interface {
	FindUser(login, password string) (*users.User, error)
}

func NewUserGetter(logger *log.Logger, userGetter UserGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.get.NewUserGetter"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var getUser request.GetUser
		var req request.Request = &getUser
		request.Decode(w, r, &req)

		log.Info("decoded request body", getUser)

		if err := validator.New().Struct(getUser); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		user, err := userGetter.FindUser(getUser.Login, getUser.Password)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to get user"))

			return
		}

		log.Info("successfully get user")
		response.ResponseOKWithData(w, r, user)
	}
}
