package v1

import (
	"context"
	"net/http"
)

type Events struct {
	prefix string
	ctx    context.Context
}

func NewEvents(prefixPath string, ctx context.Context) http.Handler {
	return &Events{
		prefix: prefixPath,
		ctx:    ctx,
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/device-update":
		// TODO: Add connection to sse event handlers (see pirgb-web for reference)

		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		http.NotFound(w, r)
	}
}
