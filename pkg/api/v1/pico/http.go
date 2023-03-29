package pico

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetDuty(addr string) (duty [4]int, err error) {
	url := fmt.Sprintf("http://%s%s", addr, PathGetDuty)
	r, err := http.Get(url)
	if err != nil {
		return duty, err
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return duty, fmt.Errorf("%s: %s", url, r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return duty, err
	} else if len(data) == 0 {
		return duty, err
	}

	for i, n := range strings.Split(strings.Trim(string(data), " \n"), " ") {
		d, err := strconv.Atoi(n)
		if err != nil {
			return duty, err
		}

		duty[i] = d
	}

	return duty, nil
}

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]int, err error) {
	url := fmt.Sprintf("http://%s%s", addr, PathGetPins)
	r, err := http.Get(url)
	if err != nil {
		return pins, err
	}

	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return pins, fmt.Errorf("%s: %s", url, r.Status)
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return pins, err
	} else if len(data) == 0 {
		return pins, fmt.Errorf("no data returned from device " + addr)
	}

	for i, n := range strings.Split(strings.Trim(string(data), " \n"), " ") {
		gp, err := strconv.Atoi(n)
		if err != nil {
			return pins, err
		}

		pins[i] = gp
	}

	return pins, nil
}

func SetDuty(addr string, rgbw [4]int) (err error) {
	// TODO: ...

	return fmt.Errorf("Under Construction")
}

func SetPins(addr string, pins [4]int) (err error) {
	// TODO: ...

	return fmt.Errorf("Under Construction")
}
