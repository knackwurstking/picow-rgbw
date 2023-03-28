package pico

const (
	// PathGetPins server pathname for getting pins in use
	PathGetPins = "/rgbw/get_pins"
)

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]int) {
	// TODO: ...

	return pins
}
