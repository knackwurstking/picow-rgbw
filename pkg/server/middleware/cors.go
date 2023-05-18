package middleware

import "net/http"

var (
	AllowOrigin = "*"
)

type CORSOptions struct {
	AllowOrigin *string
}

type CORS struct {
	handler http.Handler
	options CORSOptions
}

func NewCORS(h http.Handler) http.Handler {
	return &CORS{
		handler: h,
		options: CORSOptions{
			AllowOrigin: &AllowOrigin,
		},
	}
}

func (c *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", *c.options.AllowOrigin)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	c.handler.ServeHTTP(w, r)
}
