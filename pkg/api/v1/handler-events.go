package v1

import (
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
)

type Events struct {
	prefix string
	pico   *pico.Handler
}

func NewEvents(prefixPath string, p *pico.Handler) http.Handler {
	return &Events{
		prefix: prefixPath,
		pico:   p,
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/devices":
		conn, ok := e.pico.SSE.Add("devices", w, r)
		if !ok {
			return
		}
		defer e.pico.SSE.Close("devices", conn)
		_ = conn.Write("update", e.pico.Devices)
		<-conn.Request.Context().Done()
	case e.prefix + "/device":
		conn, ok := e.pico.SSE.Add("device", w, r)
		if !ok {
			return
		}
		defer e.pico.SSE.Close("device", conn)
		<-conn.Request.Context().Done()
	default:
		http.NotFound(w, r)
	}
}
