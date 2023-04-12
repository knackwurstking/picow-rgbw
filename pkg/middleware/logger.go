package middleware

import (
	"net/http"
	"time"

	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(h http.Handler) http.Handler {
	return &Logger{
		handler: h,
	}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: need to use a custom response writer (where i can get the status code for logging)
	defer func(startTime time.Time) {
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path += "?" + r.URL.RawQuery
		}

		log.Info.Printf(
			"[%s] - %s - \"%s\" - %s",
			r.RemoteAddr,
			r.Method,
			path,
			time.Since(startTime).String(),
		)
	}(time.Now())

	l.handler.ServeHTTP(w, r)
}
