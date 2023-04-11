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
		sse:    sse.NewHandler(), // TODO: Add sse handler to pico handler
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/devices-update":
		conn, ok := e.sse.Add("devices-update", w, r)
		if !ok {
			return
		}
		defer e.sse.Close(conn)

		<-conn.Request.Context().Done()
	case e.prefix + "/device-update":
		conn, ok := e.sse.Add("device-update", w, r)
		if !ok {
			return
		}
		defer e.sse.Close(conn)

		<-conn.Request.Context().Done()
	default:
		http.NotFound(w, r)
	}
}
