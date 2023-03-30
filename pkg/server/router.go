package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/knackwurstking/picow-rgbw-web/frontend"
	api "github.com/knackwurstking/picow-rgbw-web/pkg/api/v1"
)

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

	{ // File Server
		// TODO: change this to use frontend/frontend.go embed stuff
		mux.Routes = append(mux.Routes, &Route{
			Pattern: regexp.MustCompile("/"),
			Handler: http.FileServer(frontend.Dist()),
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
