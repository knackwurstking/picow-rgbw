package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
}

func (d *Devices) device(w http.ResponseWriter, r *http.Request, id int) {
	p, err := getPicoFromCtx(d.ctx)
	if err != nil {
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
}

func (d *Devices) putDevices(w http.ResponseWriter, r *http.Request) {
	// check for content type (json)
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		http.Error(w, http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}

	// get *pico.Handler from context
	p, err := getPicoFromCtx(d.ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	// read body data
	defer r.Body.Close()
	var data []ReqPutDevice
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
		go func(rd ReqPutDevice, statusCh chan int, doneCh chan struct{}) {
			defer func() {
				doneCh <- struct{}{}
			}()

			device := p.Get(rd.Addr)
			if device == nil {
				http.Error(w, http.StatusText(http.StatusBadRequest),
					http.StatusBadRequest)
				statusCh <- http.StatusBadRequest
				return
			}

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
			if doneCount == 0 {
				w.WriteHeader(status)
				return
			}
		}
	}
}
