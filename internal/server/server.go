package server

import (
	"net/http"

	api "github.com/knackwurstking/picow-rgbw-web/internal/api/v1"
	"github.com/knackwurstking/picow-rgbw-web/internal/middleware"
)

var (
	middlewareHandlers []func(http.Handler) http.Handler
)

// UseMiddleware for all http handlers
func UseMiddleware(h func(http.Handler) http.Handler) {
	middlewareHandlers = append(middlewareHandlers, h)
}

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()

	{ // Group: "/api/v1"
		group := "/api/v1"

		mux.HandleFunc(group+"/devices", AddMiddleware(
			api.NewDevices(group+"/devices"),
			middlewareHandlers...,
		).ServeHTTP)

		mux.HandleFunc(group+"/events", AddMiddleware(
			api.NewEvents(group+"/events"),
			middlewareHandlers...,
		).ServeHTTP)
	}

	return mux
}

func New(addr string) *http.Server {
	UseMiddleware(
		middleware.NewLogger,
	)

	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(),
	}
}
