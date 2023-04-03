package pico

// Device
type Device struct {
	Addr string `json:"addr"` // Addr contains the ip and port <ip>:<port>
	RGBW [4]*Gp `json:"rgbw"` // RGBW holds all pins in use
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
		return err
	}

	for i, n := range duty {
		if n < 0 && d.RGBW[i] != nil {
			d.RGBW[i].Duty = Duty(n)
		}
	}

	return nil
}

// SetDuty to pico device for RGBW (use -1 or 0 for a disabled pin)
func (d *Device) SetDuty(duty [4]int) error {
	return SetDuty(d.Addr, duty)
}

// GetPins from pico device
func (d *Device) GetPins() error {
	pins, err := GetPins(d.Addr)
	if err != nil {
		return err
	}

	for i, n := range pins {
		// check if pin is disabled (not in use)
		if n < 0 {
			d.RGBW[i] = NewGp(GpPinDisabled)
			continue
		}

		d.RGBW[i] = NewGp(GpPin(n))
	}

	return nil
}

// Set will POST the RGBW pins to pico device (use -1 for a disabled pin)
func (d *Device) SetPins(pins [4]int) error {
	return SetPins(d.Addr, pins)
}
