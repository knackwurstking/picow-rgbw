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
func (d *Device) GetPins() {
	for i, n := range GetPins(d.Addr) {
		// check if pin is disabled (not in use)
		if n < 0 {
			d.RGBW[i] = nil
			continue
		}

		d.RGBW[i] = NewGpPWM(n)
	}
}

// Set will POST the RGBW pins to pico device
func (d *Device) SetPins() {
	// TODO: ...
}

// GetDuty from pico device
func (d *Device) GetDuty() {
	// TODO: get duty for each pin from device
}

// SetDuty to pico device for RGBW
func (d *Device) SetDuty(rgbw [4]int) {
	// TODO: ...
}

// SetDutyR will only set the "r" duty for the pico device
func (d *Device) SetDutyR(r int) {
	// TODO: ...
}

// SetDutyG will only set the "r" duty for the pico device
func (d *Device) SetDutyG(r int) {
	// TODO: ...
}

// SetDutyB will only set the "r" duty for the pico device
func (d *Device) SetDutyB(r int) {
	// TODO: ...
}

// SetDutyW will only set the "r" duty for the pico device
func (d *Device) SetDutyW(r int) {
	// TODO: ...
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
