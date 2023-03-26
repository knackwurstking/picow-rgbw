package server

import (
	"context"
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

func NewHandler(ctx context.Context) *http.ServeMux {
	mux := http.NewServeMux()

	{ // Group: "/api/v1"
		group := "/api/v1"

		mux.HandleFunc(group+"/devices", AddMiddleware(
			api.NewDevices(group+"/devices", ctx),
			middlewareHandlers...,
		).ServeHTTP)

		mux.HandleFunc(group+"/events", AddMiddleware(
			api.NewEvents(group+"/events", ctx),
			middlewareHandlers...,
		).ServeHTTP)
	}

	return mux
}

func New(addr string, picoHandler *pico.Handler) *http.Server {
	UseMiddleware(
		middleware.NewLogger,
	)

	ctx := context.WithValue(context.Background(), "pico", picoHandler)
	return &http.Server{
		Addr:    addr,
		Handler: NewHandler(ctx),
	}
}
