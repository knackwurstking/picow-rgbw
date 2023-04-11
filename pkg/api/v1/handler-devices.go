package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gookit/slog"
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/pico"
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
		switch r.Method {
		case http.MethodGet:
			d.devices(w, r)
		case http.MethodPost:
			d.putDevices(w, r)
		default:
			http.NotFound(w, r)
		}
	case idRegEx.MatchString(path):
		switch r.Method {
		case http.MethodGet:
			id, _ := strconv.Atoi(idRegEx.FindStringSubmatch(path)[1])
			d.device(w, r, id)
		default:
			http.NotFound(w, r)
		}
	default:
		http.NotFound(w, r)
	}
}

func (d *Devices) devices(w http.ResponseWriter, r *http.Request) {
	handler := d.ctx.Value("pico").(*pico.Handler)
	if err := json.NewEncoder(w).Encode(handler.Devices); err != nil {
		slog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (d *Devices) device(w http.ResponseWriter, r *http.Request, id int) {
	handler := d.ctx.Value("pico").(*pico.Handler)
	for i, device := range handler.Devices {
		if i == id {
			if err := json.NewEncoder(w).Encode(device); err != nil {
				slog.Error(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}

func (d *Devices) putDevices(w http.ResponseWriter, r *http.Request) {
	if !hasJSONContent(w, r) {
		return
	}

	// read body data
	handler := d.ctx.Value("pico").(*pico.Handler)
	defer r.Body.Close()
	var data []RequestPostDevice
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// update devices
	statusCh := make(chan int)
	doneCh := make(chan struct{})
	doneCount := 0

	for _, reqDevice := range data {
		doneCount += 1
		go func(rd RequestPostDevice, statusCh chan int, doneCh chan struct{}) {
			defer func() {
				doneCh <- struct{}{}
			}()

			device := handler.Get(rd.Addr)
			if device == nil {
				http.Error(w, http.StatusText(http.StatusBadRequest),
					http.StatusBadRequest)
				statusCh <- http.StatusBadRequest
				return
			}

			slog.Debug(fmt.Sprintf("set rgbw (%v) for %s", rd.RGBW, device.Addr))
			if err := device.SetDuty(rd.RGBW); err != nil {
				statusCh <- http.StatusInternalServerError
				return
			}
		}(reqDevice, statusCh, doneCh)
	}

	status := http.StatusOK
	for {
		select {
		case s := <-statusCh:
			if status == http.StatusBadRequest {
				continue
			}

			status = s
		case <-doneCh:
			doneCount -= 1
			if doneCount == 0 {
				if status >= 200 && status < 300 {
					w.WriteHeader(status)
				} else {
					http.Error(w, http.StatusText(status), status)
				}
				return
			}
		}
	}
}
