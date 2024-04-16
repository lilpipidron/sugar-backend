package user

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type UserSaver interface {
	AddUser(user users.User, password string) error
}

func New(logger *log.Logger, userSaver UserSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userAdd request.AddUser
		var req request.Request = &userAdd
		request.Decode(w, r, &req)
		userAdd = (req).(request.AddUser)

		log.Info("decoded request body", userAdd)

		if err := validator.New().Struct(userAdd); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		userInfo := users.UserInfo{
			Name:              userAdd.Name,
			Birthday:          userAdd.Birthday,
			Gender:            userAdd.Gender,
			Weight:            userAdd.Weight,
			CarbohydrateRatio: userAdd.CarbohydrateRatio,
			BreadUnit:         userAdd.BreadUnit,
		}

		user := users.User{
			UserID:   -1,
			Login:    userAdd.Login,
			UserInfo: userInfo,
		}

		err := userSaver.AddUser(user, userAdd.Password)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to add user"))

			return
		}

		log.Info("added user")

		response.ResponseOK(w, r)
	}
}
