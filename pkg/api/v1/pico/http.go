package pico

const (
	// PathGetPins server pathname for getting pins in use
	PathGetPins = "/rgbw/get_pins"
	// PathGetDuty server pathname for getting the current duty for each pin
	PathGetDuty = "/rgbw/get_duty"
)

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]int, ok bool) {
	// TODO: http get request to `PathGetPins` for `addr`, parse result and return pins

	return pins, ok
}

func GetDuty(addr string) (duty [4]int, ok bool) {
	// TODO: http get request to `PathGetDuty` for `addr`, parse result and return duty

	return duty, ok
}

func SetDuty(addr string, rgbw [4]int) (ok bool) {
	// TODO: ...

	return ok
}

func SetPins(addr string, pins [4]int) (ok bool) {
	// TODO: ...

	return ok
}
