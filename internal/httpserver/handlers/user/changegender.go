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
	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type GenderChanger interface {
	ChangeGender(id int64, newGender users.Gender) error
}

func NewGenderChanger(logger *log.Logger, genderChanger GenderChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changeGender.NewGenderChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeGender request.ChangeGender
		var req request.Request = &changeGender
		request.Decode(w, r, &req)
		changeGender = (req).(request.ChangeGender)

		log.Info("decoded request body", changeGender)

		if err := validator.New().Struct(changeGender); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := genderChanger.ChangeGender(changeGender.ID, changeGender.NewGender)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change gender"))
		}

		log.Info("changed gender")

		response.ResponseOK(w, r)
	}
}
