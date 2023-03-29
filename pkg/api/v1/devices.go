package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

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
	idRegEx := regexp.MustCompile(d.prefix + `\/(\d+)`)

	switch path := r.URL.Path; {
	case path == d.prefix+"", path == d.prefix+"/":
		p, err := getPicoFromCtx(d.ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(p.Devices)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	case idRegEx.MatchString(path):
		p, err := getPicoFromCtx(d.ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := strconv.Atoi(idRegEx.FindStringSubmatch(path)[1])
		if err != nil {
			slog.Error("parsing id failed: " + err.Error()) // NOTE: should never happen
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for i, device := range p.Devices {
			if i == id {
				if err = json.NewEncoder(w).Encode(device); err != nil {
					slog.Error(err.Error())
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		http.NotFound(w, r)
	default:
		http.NotFound(w, r)
	}
}
