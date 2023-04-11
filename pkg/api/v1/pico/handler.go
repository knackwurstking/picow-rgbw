package pico

import (
	"fmt"

	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
)

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device `json:"devices"`

	// TODO: sse handler for /device-update every time devices got an update
	//		 (of device data changed)
	sse *sse.Handler `json:"-"` // sse can be nil, make sure to check first
}

// NewPico
func NewHandler(devices ...*Device) *Handler {
	return &Handler{
		Devices: devices,
	}
}

func (h *Handler) Add(device *Device) {
	for i, d := range h.Devices {
		if d.Addr == device.Addr {
			h.Devices[i] = device
		}
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
