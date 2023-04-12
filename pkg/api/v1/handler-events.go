package v1

import (
	"context"
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
)

type Events struct {
	prefix string
	ctx    context.Context
	sse    *sse.Handler
}

func NewEvents(prefixPath string, ctx context.Context) http.Handler {
	p := ctx.Value("pico").(*pico.Handler)
	p.SSE = sse.NewHandler()

	return &Events{
		prefix: prefixPath,
		ctx:    ctx,
		sse:    p.SSE,
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/devices":
		conn, ok := e.sse.Add("devices", w, r)
		if !ok {
			return
		}
		defer e.sse.Close("devices", conn)
		<-conn.Request.Context().Done()
	case e.prefix + "/device":
		conn, ok := e.sse.Add("device", w, r)
		if !ok {
			return
		}
		defer e.sse.Close("device", conn)
		<-conn.Request.Context().Done()
	default:
		http.NotFound(w, r)
	}
}
