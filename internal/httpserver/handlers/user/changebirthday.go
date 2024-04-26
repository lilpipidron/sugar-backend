package user

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/request"
	"github.com/lilpipidron/sugar-backend/internal/httpserver/response"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
)

type BirthdayChanger interface {
	ChangeBirthday(id int64, newBirthday time.Time) error
}

func NewBirthdayChanger(logger *log.Logger, birthdayChanger BirthdayChanger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.changeBirthday.NewBirthdayChanger"

		logger = log.With(
			"op: "+op,
			"request_id"+middleware.GetReqID(r.Context()),
		)

		var changeBirthday request.ChangeBirthday
		var req request.Request = &changeBirthday
		request.Decode(w, r, &req)
		changeBirthday = (req).(request.ChangeBirthday)

		log.Info("decoded request body", changeBirthday)

		if err := validator.New().Struct(changeBirthday); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		err := birthdayChanger.ChangeBirthday(changeBirthday.ID, changeBirthday.NewBirthday)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to change birthday"))

			return
		}

		log.Info("changed birthday")

		response.ResponseOK(w, r)
	}
}
