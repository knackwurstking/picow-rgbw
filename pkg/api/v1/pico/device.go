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

// GetDuty from pico device
func (d *Device) GetDuty() error {
	duty, err := GetDuty(d.Addr)
	if err != nil {
		if IsUrlError(err) {
			d.Offline = true
		}

		return err
	}

	for i, n := range duty {
		if d.RGBW[i] == nil {
			d.RGBW[i] = NewGp(GpPinDisabled)
		}

		if n < DutyMin {
			d.RGBW[i].Duty = Duty(DutyMin)
		} else if n > DutyMax {
			d.RGBW[i].Duty = Duty(DutyMax)
		} else {
			d.RGBW[i].Duty = Duty(n)
		}
	}

	d.Offline = false

	d.update()
	return nil
}

// SetDuty to pico device for RGBW (use -1 or 0 for a disabled pin)
func (d *Device) SetDuty(duty [4]Duty) error {
	err := SetDuty(d.Addr, duty)
	if err == nil {
		for i, gp := range d.RGBW {
			gp.Duty = duty[i]
		}
	}

	if IsUrlError(err) {
		d.Offline = true
	} else {
		d.Offline = false
	}

	d.update()
	return err
}

// GetPins from pico device
func (d *Device) GetPins() error {
	pins, err := GetPins(d.Addr)
	if err != nil {
		if IsUrlError(err) {
			d.Offline = true
		}

		return err
	}

	for i, n := range pins {
		if n >= GpPinMin && n <= GpPinMax {
			d.RGBW[i] = NewGp(GpPin(n))
		} else {
			d.RGBW[i] = NewGp(GpPinDisabled)
		}
	}

	d.Offline = false

	d.update()
	return nil
}

// Set will POST the RGBW pins to pico device (use -1 for a disabled pin)
func (d *Device) SetPins(pins [4]GpPin) error {
	err := SetPins(d.Addr, pins)
	if err == nil {
		for i, gp := range d.RGBW {
			gp.Nr = pins[i]
		}
	}

	if IsUrlError(err) {
		d.Offline = true
	} else {
		d.Offline = false
	}

	d.update()
	return err
}

func (d *Device) update() {
	if d.SSE != nil {
		d.SSE.Dispatch("device", "update", d)
	} else {
		log.Debug.Println("SSE handler missing", d)
	}
}
