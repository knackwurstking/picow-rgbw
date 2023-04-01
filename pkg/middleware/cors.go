package middleware

import "net/http"

type CORS struct {
	handler http.Handler
}

func NewCORS(h http.Handler) http.Handler {
	return &CORS{
		handler: h,
	}
}

func (c *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: ...

	c.handler.ServeHTTP(w, r)
}
