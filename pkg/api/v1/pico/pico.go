package pico

// GpPWM
type GpPWM struct {
	Nr   int `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
	Duty int `json:"duty"` // Duty cycle (goes from 0-100)
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

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device `json:"devices"`
}

// NewPico
func NewPico(devices ...*Device) *Handler {
	return &Handler{
		Devices: devices,
	}
}

// Scan for pico devices
func (h *Handler) Scan() (devices []*Device, err error) {
	// TODO: Scan the local network for pico devices

	return devices, nil
}
