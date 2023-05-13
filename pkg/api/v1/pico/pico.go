package pico

import "strconv"

const (
	DutyMin = Duty(0)
	DutyMax = Duty(100)

	TCPGetColorCommand = "rgbw color get;"

	TCPGetColorSize = 255
)

type Duty int

func (d *Duty) String() string {
	return strconv.Itoa(int(*d))
}
