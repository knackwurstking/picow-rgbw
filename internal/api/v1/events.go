package v1

import "net/http"

type Events struct {
	prefix string
}

func NewEvents(prefixPath string) http.Handler {
	return &Events{
		prefix: prefixPath,
	}
}

func (e *Events) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case e.prefix + "/device-update":
		// TODO: Add connection to sse event handlers (see pirgb-web for reference)

		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
