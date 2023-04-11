package v1

import (
	"context"
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
)

type Events struct {
	prefix string
	ctx    context.Context
	sse    *sse.Handler
}

func NewEvents(prefixPath string, ctx context.Context) http.Handler {
	return &Events{
		prefix: prefixPath,
		ctx:    ctx,
		sse:    sse.NewHandler(),
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/device-update":
		// TODO: Add connection to sse event handlers (see pirgb-web for reference)

		http.Error(w, http.StatusText(http.StatusServiceUnavailable),
			http.StatusServiceUnavailable)
	default:
		http.NotFound(w, r)
	}
}
