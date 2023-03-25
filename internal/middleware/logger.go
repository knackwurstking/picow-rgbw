package middleware

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
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
	defer func(startTime time.Time) {
		path := r.URL.Path
		if r.URL.RawQuery != "" {
			path += "?" + r.URL.RawQuery
		}

		slog.Info(
			fmt.Sprintf("[%s] - %s - \"%s\" - %s",
				r.RemoteAddr,
				r.Method,
				path,
				time.Since(startTime).String(),
			),
		)
	}(time.Now())

	l.handler.ServeHTTP(w, r)
}