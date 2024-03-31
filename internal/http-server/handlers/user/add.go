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
	CarbohydrateRatio int          `json:carbohydrate-ratio"`
	BreadUnit         int          `json:"bread-unit`
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

		var req Request

		err := render.DecodeJSON(r.Body, &req)
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

		log.Info("request body decoded", req)

		if err != validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", err)

			render.JSON(w, r, resp.Error(validateErr.Error()))

			return
		}

		usrInfo := users.UserInfo{
			req.Name,
			req.Birthday,
			req.Gender,
			req.Weight,
			req.CarbohydrateRatio,
			req.BreadUnit,
		}

		usr := users.User{
			-1,
			req.Login,
			usrInfo,
		}

		err = userSaver.AddUser(usr, req.Password)
		if err != nil {
			log.Error(err)

			render.JSON(w, r, resp.Error("failed add user"))

			return
		}

		log.Info("user added")

		responseOK(w, r)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}
