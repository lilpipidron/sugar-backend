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
	"time"
)

type AllInfoChanger interface {
	ChangeAllInfo(userID int64, newName string, newBirthday time.Time, newGender users.Gender, newWeight int, newCarbohydrate int, newBreadUnit int, newHeight int) error
}

func NewUserInfoChanger(logger *log.Logger, allInfoChanger AllInfoChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changeinfo.NewUserInfoChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeUserInfo request.ChangeUserInfo
		var req request.Request = &changeUserInfo
		request.Decode(w, r, &req)

		log.Info("decoded request body", changeUserInfo)

		if err := validator.New().Struct(changeUserInfo); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := allInfoChanger.ChangeAllInfo(changeUserInfo.ID, changeUserInfo.NewName, changeUserInfo.NewBirthday, changeUserInfo.NewGender,
			changeUserInfo.NewWeight, changeUserInfo.NewCarbohydrate, changeUserInfo.NewBreadUnit, changeUserInfo.NewHeight)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change birthday"))

			return
		}

		log.Info("changed birthday")

		response.ResponseOK(w, r)
	}
}
