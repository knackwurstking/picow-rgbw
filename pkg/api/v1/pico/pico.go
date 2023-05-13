package pico

import "strconv"

const (
	DutyMin = Duty(0)
	DutyMax = Duty(100)

	TCPGetColorCommand = "rgbw color get;"
	TCPSetColorCommand = "rgbw color set;"
	TCPGetGpCommand    = "rgbw gp get;"
	TCPSetGpCommand    = "rgbw gp set;"

	TCPGetColorReadSize = 255
	TCPSetColorReadSize = -1
	TCPGetGpReadSize    = 255
	TCPSetGpReadSize    = -1
)

type Duty int

func (d *Duty) String() string {
	return strconv.Itoa(int(*d))
}
