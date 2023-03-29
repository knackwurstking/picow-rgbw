package pico

import "net/http"

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
