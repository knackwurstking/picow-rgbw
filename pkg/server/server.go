package server

import (
	"net/http"

	api "github.com/knackwurstking/picow-rgbw-web/pkg/api/v1"
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/middleware"
)

var (
	middlewareHandlers []func(http.Handler) http.Handler
)

// UseMiddleware for all http handlers
func UseMiddleware(h func(http.Handler) http.Handler) {
	middlewareHandlers = append(middlewareHandlers, h)
}

func NewHandler(handler *pico.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	{ // Group: "/api/v1"
		group := "/api/v1"

		mux.HandleFunc(group+"/devices", AddMiddleware(
			api.NewDevices(group+"/devices", handler),
			middlewareHandlers...,
		).ServeHTTP)

		mux.HandleFunc(group+"/events", AddMiddleware(
			api.NewEvents(group+"/events"),
			middlewareHandlers...,
		).ServeHTTP)
	}

	return mux
}

func New(addr string, picoHandler *pico.Handler) *http.Server {
	// TODO: Using context instead of pico handler here (key: "pico", value: "*pico.Handler")
	UseMiddleware(
		middleware.NewLogger,
	)

	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(picoHandler),
	}
}
