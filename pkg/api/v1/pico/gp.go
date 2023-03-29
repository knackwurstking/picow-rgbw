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
