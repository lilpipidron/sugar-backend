package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
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
			)

			t1 := time.Now()

			defer func() {
				entry.Info("request completed",
					fmt.Sprintf("status %d", r.Response.StatusCode),
					fmt.Sprintf("content length %d", r.Response.ContentLength),
					fmt.Sprintf(`duration %s`, time.Since(t1).String()),
				)
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
