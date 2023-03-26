package server

import "net/http"

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
