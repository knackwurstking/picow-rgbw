package v1

import (
	"encoding/json"
	"net/http"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
)

type PicoW struct {
	prefix string
	pico   *pico.Handler
}

func NewPicoW(prefixPath string, p *pico.Handler) http.Handler {
	return &PicoW{
		prefix: prefixPath,
		pico:   p,
	}
}

func (p *PicoW) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.postPico(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *PicoW) postPico(w http.ResponseWriter, r *http.Request) {
	if !hasJSONContent(w, r) {
		return
	}

	defer r.Body.Close()
	var data pico.Device
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil || data.Addr == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	device := p.pico.Get(data.Addr)
	update := device != nil
	if !update {
		device = &data
	}

	go func() {
		err := device.GetGpPins()
		if err != nil {
			log.Warn.Printf("Get pins: %s", err.Error())
			return
		}

		_ = device.GetColor()

		if update {
			p.pico.Add(device)
		}
	}()

	w.WriteHeader(http.StatusOK)
}
