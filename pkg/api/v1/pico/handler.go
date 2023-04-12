package pico

import (
	"fmt"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
)

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device    `json:"devices"`
	SSE     *sse.Handler `json:"-"` // sse can be nil, make sure to check first
}

// NewPico
func NewHandler(devices ...*Device) *Handler {
	return &Handler{
		Devices: devices,
	}
}

func (h *Handler) Add(device *Device) {
	device.SSE = h.SSE

	for i, d := range h.Devices {
		if d.Addr == device.Addr {
			h.Devices[i] = device
			if h.SSE != nil {
				h.update()
			}
			return
		}
	}

	h.Devices = append(h.Devices, device)
	if h.SSE != nil {
		h.update()
	}
}

func (h *Handler) Get(addr string) *Device {
	for _, d := range h.Devices {
		if d.Addr == addr {
			return d
		}
	}

	return nil
}

// Scan for pico devices (ex. r: 192.168.178.0, 192.168.0.0)
func (h *Handler) Scan(ipRange string) (devices []*Device, err error) {
	// TODO: Scan ip range...
	// ...

	return devices, fmt.Errorf("Scanner Not Implemented!")
}

func (h *Handler) update() {
	if h.SSE != nil {
		h.SSE.Dispatch("devices", "update", h)
	}
}
