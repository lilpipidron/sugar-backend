package user

import (
	"github.com/lilpipidron/sugar-backend/internal/crypt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type UserSaver interface {
	AddNewUser(user users.User, password string) error
}

func NewUserSaver(logger *log.Logger, userSaver UserSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.add.NewUserSaver"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var userAdd request.AddUser
		var req request.Request = &userAdd
		request.Decode(w, r, &req)

		log.Info("decoded request body", userAdd)

		if err := validator.New().Struct(userAdd); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		var err error = nil
		userAdd.Password, err = crypt.HashPassword(userAdd.Password)

		if err != nil {
			log.Error("failed to hash password", "error", err)

			render.Status(r, http.StatusInternalServerError)

			render.JSON(w, r, resp.Error(err.Error()))
		}

		userInfo := users.UserInfo{
			Name:              userAdd.Name,
			Birthday:          userAdd.Birthday,
			Gender:            userAdd.Gender,
			Weight:            userAdd.Weight,
			CarbohydrateRatio: userAdd.CarbohydrateRatio,
			BreadUnit:         userAdd.BreadUnit,
			Height:            userAdd.Height,
		}

		user := users.User{
			UserID:   -1,
			Login:    userAdd.Login,
			UserInfo: userInfo,
		}

		err = userSaver.AddNewUser(user, userAdd.Password)
		if err != nil {
			log.Error(err)

			render.Status(r, http.StatusBadRequest)

			render.JSON(w, r, resp.Error("failed to add user"))

			return
		}

		log.Info("added user")

		response.ResponseOK(w, r)
	}
}
