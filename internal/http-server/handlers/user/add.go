package user

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
	"github.com/lilpipidron/sugar-backend/internal/models/users"
)

type Request struct {
	Login             string       `json:"login"`
	Password          string       `json:"password"`
	Name              string       `json:"name"`
	Birthday          time.Time    `json:"birthday"`
	Gender            users.Gender `json:"gender"`
	Weight            int          `json:"weight"`
	CarbohydrateRatio int          `json:"carbohydrate-ratio"`
	BreadUnit         int          `json:"bread-unit"`
}

type Response struct {
	resp.Response
}
type UserSaver interface {
	AddUser(user users.User, password string) error
}

func New(logger *log.Logger, userSaver UserSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user.add.New"
		logger = log.With(
			"op: "+op,
			"request_id: "+middleware.GetReqID(r.Context()),
		)

		var request Request

		err := render.DecodeJSON(r.Body, &request)
		if errors.Is(err, io.EOF) {
			log.Errorf("request body is empty")

			render.JSON(w, r, resp.Error("empty request"))

			return
		}

		if err != nil {
			log.Error("failed to decode request body", err)

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("decoded request body", request)

		if err = validator.New().Struct(request); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		userInfo := users.UserInfo{
			Name:              request.Name,
			Birthday:          request.Birthday,
			Gender:            request.Gender,
			Weight:            request.Weight,
			CarbohydrateRatio: request.CarbohydrateRatio,
			BreadUnit:         request.BreadUnit,
		}

		user := users.User{
			UserID:   -1,
			Login:    request.Login,
			UserInfo: userInfo,
		}

		err = userSaver.AddUser(user, request.Password)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed to add user"))

			return
		}

		log.Info("added user")

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
