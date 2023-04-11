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
		handler, ok := getHandler(w, e.ctx)
		if !ok {
			return
		}

		conn, ok := e.sse.Add(w, r, handler)
		if !ok {
			return
		}
		defer e.sse.Close(conn)

		<-conn.Request.Context().Done()
	default:
		http.NotFound(w, r)
	}
}
