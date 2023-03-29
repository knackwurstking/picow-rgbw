package pico

import (
	"fmt"
	"net/http"
)

const (
	// PathGetPins server pathname for getting pins in use
	PathGetPins = "/rgbw/get_pins"
	// PathGetDuty server pathname for getting the current duty for each pin
	PathGetDuty = "/rgbw/get_duty"

	// StatusOffline device not reachable (offline?)
	StatusOffline = 0
	// StatusOK
	StatusOK = http.StatusOK
	// StatusBadRequest
	StatusBadRequest = http.StatusBadRequest
	// StatusNotFound
	StatusNotFound = http.StatusNotFound
	// StatusInternalServerError
	StatusInternalServerError = http.StatusInternalServerError
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
