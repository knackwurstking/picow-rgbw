package pico

import (
	"fmt"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
)

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device   `json:"devices"`
	SSE     sse.Handler `json:"-"` // sse can be nil, make sure to check first
}

// NewPico
func NewHandler() *Handler {
	h := &Handler{
		Devices: make([]*Device, 0),
		SSE: sse.Handler{
			Connections: sse.NewConnections(),
		},
	}
	return h
}

func (h *Handler) Add(device *Device) {
	device.SSE = &h.SSE

	for i, d := range h.Devices {
		if d.Addr == device.Addr {
			h.Devices[i] = device
			h.update()
			return
		}
	}

	h.Devices = append(h.Devices, device)
	h.update()
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
	h.SSE.Dispatch("devices", "update", h)
}
