package pico

import "strconv"

const (
	DutyMin = Duty(0)
	DutyMax = Duty(100)

	TCPGetColorCommand  = "rgbw color get;"
	TCPSetColorCommand  = "rgbw color set %d %d %d %d;"
	TCPGetGpPinsCommand = "rgbw gp get;"
	TCPSetGpPinsCommand = "rgbw gp set %d %d %d %d;"

	TCPGetColorReadSize  = 255
	TCPSetColorReadSize  = -1
	TCPGetGpPinsReadSize = 255
	TCPSetGpPinsReadSize = -1
)

type Duty int

func (d *Duty) String() string {
	return strconv.Itoa(int(*d))
}
