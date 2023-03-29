package pico

import (
	"fmt"
	"net/http"
)

// StatusText
func StatusText(status int) string {
	switch status {
	case 0:
		return fmt.Sprintf("%d OFFLINE", status)
	default:
		return http.StatusText(status)
	}
}

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]int, status int) {
	// TODO: http get request to `PathGetPins` for `addr`, parse result and return pins

	return pins, StatusOffline
}

func GetDuty(addr string) (duty [4]int, status int) {
	// TODO: http get request to `PathGetDuty` for `addr`, parse result and return duty

	return duty, StatusOffline
}

func SetDuty(addr string, rgbw [4]int) (status int) {
	// TODO: ...

	return StatusOffline
}

func SetPins(addr string, pins [4]int) (status int) {
	// TODO: ...

	return StatusOffline
}
