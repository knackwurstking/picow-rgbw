package pico

import "strconv"

const (
	GpPinDisabled = GpPin(-1)
	DutyMin       = Duty(0)
	DutyMax       = Duty(100)
)

type GpPin int

func (p *GpPin) String() string {
	return strconv.Itoa(int(*p))
}

type Duty int

func (d *Duty) String() string {
	return strconv.Itoa(int(*d))
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
