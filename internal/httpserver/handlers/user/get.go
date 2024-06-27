package user

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/crypt"
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
			"request_id: "+middleware.GetReqID(r.Context()),
		)

		login := r.URL.Query().Get("login")
		password := r.URL.Query().Get("password")

		password = crypt.HashPassword(password)

		getUser := request.GetUser{
			Login:    login,
			Password: password,
		}

		log.Info("decoded query parameters", getUser)

		if err := validator.New().Struct(getUser); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		user, err := userGetter.FindUser(getUser.Login, getUser.Password)
		if err != nil {
			log.Error(err)
			if err.Error() == "user not found" {
				render.Status(r, http.StatusUnauthorized)
			} else {
				render.Status(r, http.StatusInternalServerError)
			}
			render.JSON(w, r, resp.Error("failed to get user"))

			return
		}

		log.Info("successfully get user")
		response.ResponseOKWithData(w, r, user)
	}
}
