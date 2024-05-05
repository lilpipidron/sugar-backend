package response

import (
	"net/http"

	"github.com/go-chi/render"
	resp "github.com/lilpipidron/sugar-backend/internal/lib/api/response"
)

type Response struct {
	resp.Response
	Data interface{}
}

func ResponseOK(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
	})
}

func ResponseOKWithData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Data:     data,
	})
}
