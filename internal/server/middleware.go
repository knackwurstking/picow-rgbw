package server

import "net/http"

type LoggerMiddleware struct {
	handler http.Handler
}

func NewLoggerMiddleware(h http.Handler) http.Handler {
	return &LoggerMiddleware{
		handler: h,
	}
}

func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)
}

func AddMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for _, mH := range reverseMiddleware(m) {
		h = mH(h)
	}
	return h
}

func reverseMiddleware(h []func(http.Handler) http.Handler) []func(http.Handler) http.Handler {
	if len(h) == 0 {
		return h
	}
	return append(reverseMiddleware(h[1:]), h[0])
}
