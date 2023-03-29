package pico

import (
	"fmt"
)

// GpPWM
type GpPWM struct {
	Nr   int `json:"nr"` // Nr of gpio pin in use (gp0 - gp28)
	Duty int `json:"-"`  // Duty cycle (goes from 0-100)
}

// NewGpPWM
func NewGpPWM(nr int) *GpPWM {
	return &GpPWM{
		Nr: nr,
	}
}

// Device
type Device struct {
	ID   int       `json:"id"`   // ID is unique
	Addr string    `json:"addr"` // Addr contains the ip and port <ip>:<port>
	RGBW [4]*GpPWM `json:"rgbw"` // RGBW holds all pins in use
}

// NewDevice
func NewDevice(id int, addr string, rgbw [4]*GpPWM) *Device {
	return &Device{
		ID:   id,
		Addr: addr,
		RGBW: rgbw,
	}
}

// GetPins from pico device
func (d *Device) GetPins() error {
	if pins, ok := GetPins(d.Addr); ok {
		for i, n := range pins {
			// check if pin is disabled (not in use)
			if n < 0 {
				d.RGBW[i] = nil
				continue
			}

			d.RGBW[i] = NewGpPWM(n)
		}

		return nil
	}

	return fmt.Errorf("get pins from %s failed", d.Addr)
}

// Set will POST the RGBW pins to pico device (use -1 for a disabled pin)
func (d *Device) SetPins(pins [4]int) error {
	if ok := SetPins(d.Addr, pins); !ok {
		return fmt.Errorf(
			"set pins for %s failed (r=%d, g=%d, b=%d, w=%d)",
			d.Addr, pins[0], pins[1], pins[2], pins[3],
		)
	}

	return nil
}

// GetDuty from pico device
func (d *Device) GetDuty() error {
	if duty, ok := GetDuty(d.Addr); ok {
		for i, n := range duty {
			if n < 0 && d.RGBW[i] != nil {
				d.RGBW[i].Duty = n
			}
		}

		return nil
	}

	return fmt.Errorf("get duty from %s failed", d.Addr)
}

// SetDuty to pico device for RGBW (use -1 or 0 for a disabled pin)
func (d *Device) SetDuty(duty [4]int) error {
	if ok := SetDuty(d.Addr, duty); !ok {
		return fmt.Errorf(
			"set duty for %s failed (r=%d, g=%d, b=%d, w=%d)",
			d.Addr, duty[0], duty[1], duty[2], duty[3],
		)
	}

	return nil
}

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device `json:"devices"`
}

// NewPico
func NewHandler(devices ...*Device) *Handler {
	return &Handler{
		Devices: devices,
	}
}

// Scan for pico devices (ex. r: 192.168.178.0, 192.168.0.0)
func (h *Handler) Scan(ipRange string) (devices []*Device, err error) {
	// TODO: Scan ip range...
	// ...

	return devices, fmt.Errorf("Scanner Not Implemented!")
}
