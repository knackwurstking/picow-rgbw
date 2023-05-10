package pico

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetDuty(addr string) (duty [4]Duty, err error) {
	url := fmt.Sprintf("http://%s%s", addr, PathGetDuty())

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

		duty[i] = Duty(d)
	}

	return duty, nil
}

// GetPins returns a list with rgbw pins in use (-1 if not in use)
func GetPins(addr string) (pins [4]GpPin, err error) {
	url := fmt.Sprintf("http://%s%s", addr, PathGetPins())

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

		pins[i] = GpPin(gp)
	}

	return pins, nil
}

func SetDuty(addr string, rgbw [4]Duty) (err error) {
	url := fmt.Sprintf("http://%s%s",
		addr, PathSetDuty(rgbw[0], rgbw[1], rgbw[2], rgbw[3]))

	r, err := http.Post(url, "text/text", nil)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", url, r.Status)
	}

	return nil
}

func SetPins(addr string, pins [4]GpPin) (err error) {
	url := fmt.Sprintf("http://%s%s",
		addr, PathSetPins(pins[0], pins[1], pins[2], pins[3]))

	r, err := http.Post(url, "text/text", nil)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", url, r.Status)
	}

	return nil
}

func IsUrlError(err error) bool {
	switch err.(type) {
	case *url.Error:
		return true
	default:
		return false
	}
}
