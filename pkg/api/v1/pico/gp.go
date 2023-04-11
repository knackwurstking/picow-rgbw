package pico

import "strconv"

const (
	GpPinMin      = GpPin(0)
	GpPinMax      = GpPin(28)
	GpPinDisabled = GpPin(-1)
)

type GpPin int

func (p *GpPin) String() string {
	return strconv.Itoa(int(*p))
}

// GpPWM
type Gp struct {
	Nr   GpPin `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
	Duty Duty  `json:"duty"` // Duty cycle (goes from 0-100)
}

// NewGpPWM
func NewGp(pin GpPin) *Gp {
	return &Gp{
		Nr: pin,
	}
}

func (gp *Gp) IsDisabled() bool {
	return gp.Nr == GpPinDisabled
}
