package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gookit/slog"
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
)

type Pico struct {
	prefix string
	ctx    context.Context
}

func NewPico(prefixPath string, ctx context.Context) http.Handler {
	return &Pico{
		prefix: prefixPath,
		ctx:    ctx,
	}
}

func (p *Pico) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.postPico(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Pico) postPico(w http.ResponseWriter, r *http.Request) {
	if !hasJSONContent(w, r) {
		return
	}

	handler, ok := getHandler(w, p.ctx)
	if !ok {
		return
	}

	defer r.Body.Close()
	var data pico.Device
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil || data.Addr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	device := handler.Get(data.Addr)
	update := device != nil
	if !update {
		device = &data
	}

	go func() {
		err := device.GetPins()
		if err != nil {
			slog.Warn("get pins: " + err.Error())
			return
		}

		_ = device.GetDuty()

		if update {
			handler.Devices = append(handler.Devices, device)
		}
	}()

	w.WriteHeader(http.StatusOK)
}
