package pico

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]int, err error) {
	r, err := http.Get(fmt.Sprintf("http://%s/%s", addr, PathGetPins))
	if err != nil {
		return pins, err
	}

	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return pins, err
	}

	for i, n := range strings.Split(strings.Trim(string(data), " \n"), " ") {
		gp, err := strconv.Atoi(n)
		if err != nil {
			return pins, err
		}

		if gp < 0 {
			// pin disabled
			pins[i] = gp
		}
	}

	return pins, err
}

func GetDuty(addr string) (duty [4]int, err error) {
	// TODO: http get request to `PathGetDuty` for `addr`, parse result and return duty

	return duty, fmt.Errorf("Under Construction")
}

func SetDuty(addr string, rgbw [4]int) (err error) {
	// TODO: ...

	return fmt.Errorf("Under Construction")
}

func SetPins(addr string, pins [4]int) (err error) {
	// TODO: ...

	return fmt.Errorf("Under Construction")
}
