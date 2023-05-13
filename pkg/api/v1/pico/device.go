package pico

import (
	"github.com/knackwurstking/picow-rgbw-web/pkg/api/v1/sse"
	"github.com/knackwurstking/picow-rgbw-web/pkg/log"
)

// Device
type Device struct {
	Addr    string       `json:"addr"`    // Addr contains the ip and port <ip>:<port>
	Offline bool         `json:"offline"` // Offline
	RGBW    [4]*Gp       `json:"rgbw"`    // RGBW holds all pins in use
	SSE     *sse.Handler `json:"-"`
}

// NewDevice
func NewDevice(addr string, rgbw [4]*Gp) *Device {
	return &Device{
		Addr: addr,
		RGBW: rgbw,
	}
}

// GetColor from picow device
func (device *Device) GetColor() error {
	color, err := GetColor(device.Addr)
	device.Offline = IsOffline(err)
	if err != nil {
		return err
	}

	for i, d := range color {
		if device.RGBW[i] == nil {
			device.RGBW[i] = NewGp(GpPinDisabled)
		}

		if d < DutyMin {
			device.RGBW[i].Duty = Duty(DutyMin)
		} else if d > DutyMax {
			device.RGBW[i].Duty = Duty(DutyMax)
		} else {
			device.RGBW[i].Duty = Duty(d)
		}
	}

	device.update()
	return nil
}

// SetDuty to picow device for RGBW (use -1 or 0 for a disabled pin)
func (device *Device) SetColor(duty [4]Duty) error {
	err := SetColor(device.Addr, duty)
	device.Offline = IsOffline(err)
	if err != nil {
		return err
	}

	for i, gp := range device.RGBW {
		gp.Duty = duty[i]
	}

	device.update()
	return err
}

// GetPins from pico device
func (device *Device) GetGpPins() error {
	gpPins, err := GetGpPins(device.Addr)
	device.Offline = IsOffline(err)
	if err != nil {
		return err
	}

	for i, p := range gpPins {
		if p >= GpPinMin && p <= GpPinMax {
			device.RGBW[i] = NewGp(p)
		} else {
			device.RGBW[i] = NewGp(GpPinDisabled)
		}
	}

	device.update()
	return nil
}

// Set will POST the RGBW pins to pico device (use -1 for a disabled pin)
func (device *Device) SetGpPins(pins [4]GpPin) error {
	err := SetGpPins(device.Addr, pins)
	device.Offline = IsOffline(err)
	device.Offline = IsOffline(err)
	if err != nil {
		return err
	}

	for i, gp := range device.RGBW {
		gp.Nr = pins[i]
	}

	device.update()
	return err
}

func (d *Device) update() {
	if d.SSE != nil {
		d.SSE.Dispatch("device", "update", d)
	} else {
		log.Debug.Println("SSE handler missing", d)
	}
}
