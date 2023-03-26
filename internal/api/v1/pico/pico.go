package pico

// GpPWM
type GpPWM struct {
	Nr   int `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
	Duty int `json:"duty"` // Duty cycle (goes from 0-100)
}

// Device
type Device struct {
	ID   int       `json:"id"`   // ID is unique
	Addr string    `json:"addr"` // Addr contains the ip and port <ip>:<port>
	RGBW [4]*GpPWM `json:"rgbw"` // RGBW holds all pins in use
}
