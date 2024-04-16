package request

import (
	"errors"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/go-chi/render"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
)

type Request interface{}

func Decode(w http.ResponseWriter, r *http.Request, request *Request) {
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
}
