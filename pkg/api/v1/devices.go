package v1

import (
	"context"
	"net/http"
	"regexp"

	"golang.org/x/exp/slog"
)

type Devices struct {
	prefix string
	ctx    context.Context
}

func NewDevices(prefixPath string, ctx context.Context) http.Handler {
	return &Devices{
		prefix: prefixPath,
		ctx:    ctx,
	}
}

func (d *Devices) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idRegEx := regexp.MustCompile(d.prefix + `\/\d+`)

	switch path := r.URL.Path; {
	case path == d.prefix+"", path == d.prefix+"/":
		_, err := getPicoFromCtx(d.ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// TODO: Return all pico devices as json data

		w.WriteHeader(http.StatusServiceUnavailable)
	case idRegEx.MatchString(path):
		_, err := getPicoFromCtx(d.ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO: Regex for :id (type int) and return the pico device data if available
		slog.Debug(":id regex match: " + path) // TODO: remove this

		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
