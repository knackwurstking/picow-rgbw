package v1

import (
	"net/http"
	"regexp"

	"golang.org/x/exp/slog"
)

type Devices struct {
	prefix string
}

func NewDevices(prefixPath string) http.Handler {
	return &Devices{
		prefix: prefixPath,
	}
}

func (d *Devices) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idRegEx := regexp.MustCompile(d.prefix + `\/\d+`)

	switch path := r.URL.Path; {
	case path == d.prefix+"", path == d.prefix+"/":
		// TODO: Return all pico devices as json data

		w.WriteHeader(http.StatusServiceUnavailable)
	case idRegEx.MatchString(path):
		// TODO: Regex for :id (type int) and return the pico device data if available
		slog.Debug(":id regex match: " + path) // TODO: remove this

		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
