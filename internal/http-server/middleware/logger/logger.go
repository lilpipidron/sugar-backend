package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5/middleware"
)

func New(logger *log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger = log.With(
			"component",
			"middleware/logger",
		)
		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				"method "+r.Method,
				"path "+r.URL.Path,
				"remote_addr"+r.RemoteAddr,
				"user_agent"+r.UserAgent(),
				"request_id"+middleware.GetReqID(r.Context()),
			)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()

			defer func() {
				entry.Info("request completed",
					fmt.Sprintf("status %w", ww.Status()),
					fmt.Sprintf("bytes %d", ww.BytesWritten()),
					fmt.Sprintf(`duration %s`, time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
