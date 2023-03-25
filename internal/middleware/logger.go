package middleware

import "net/http"

type Logger struct {
	handler http.Handler
}

func NewLogger(h http.Handler) http.Handler {
	return &Logger{
		handler: h,
	}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: defer log the current path

	l.handler.ServeHTTP(w, r)
}
