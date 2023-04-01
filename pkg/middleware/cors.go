package middleware

import "net/http"

type CORSOptions struct {
	AllowOrigin string
}

func NewCORSOptions() *CORSOptions {
	return &CORSOptions{
		AllowOrigin: "*",
	}
}

type CORS struct {
	handler http.Handler
	options *CORSOptions
}

func NewCORS(h http.Handler, options *CORSOptions) http.Handler {
	if options == nil {
		options = NewCORSOptions()
	}

	return &CORS{
		handler: h,
		options: options,
	}
}

func (c *CORS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Access-Control-Allow-Origin", c.options.AllowOrigin)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	c.handler.ServeHTTP(w, r)
}
