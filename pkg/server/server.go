package server

import (
	"context"
	"net/http"
	"regexp"

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

type Route struct {
	Pattern *regexp.Regexp
	Handler http.Handler
}

type RegExHandler struct {
	Routes []*Route
}

func NewRegExHandler(ctx context.Context) http.Handler {
	mux := &RegExHandler{
		Routes: make([]*Route, 0),
	}

	{ // Group: "/api/v1"
		group := "/api/v1"

		mux.Routes = append(mux.Routes, &Route{
			Pattern: regexp.MustCompile(group + "/devices"),
			Handler: AddMiddlewareToHandler(
				api.NewDevices(group+"/devices", ctx),
				middlewareHandlers...,
			),
		})

		mux.Routes = append(mux.Routes, &Route{
			Pattern: regexp.MustCompile(group + "/events"),
			Handler: AddMiddlewareToHandler(
				api.NewEvents(group+"/events", ctx),
				middlewareHandlers...,
			),
		})
	}

	return mux
}

func (h *RegExHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.Routes {
		if route.Pattern.MatchString(r.URL.Path) {
			route.Handler.ServeHTTP(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func New(addr string, picoHandler *pico.Handler) *http.Server {
	UseMiddleware(
		middleware.NewLogger,
	)

	// TODO: use custom type as key (replace "pico")
	ctx := context.WithValue(context.Background(), "pico", picoHandler)
	return &http.Server{
		Addr:    addr,
		Handler: NewRegExHandler(ctx),
	}
}
