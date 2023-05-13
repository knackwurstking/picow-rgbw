package pico

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type DialError struct {
	err error
}

func (d *DialError) Error() string {
	return d.err.Error()
}

func IsOffline(err error) bool {
	switch err.(type) {
	case *DialError:
		return true
	default:
		return false
	}
}

// GetColor from picow device
func GetColor(addr string) (color [4]Duty, err error) {
	// dial to picow device
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return color, &DialError{err}
	}

	// send command
	n, err := conn.Write([]byte(TCPGetColorCommand))
	if err != nil {
		return color, err
	} else if n == 0 {
		return color, fmt.Errorf("no data written to %s", addr)
	}

	// read response into data
	data := make([]byte, TCPGetColorSize)
	n, err = conn.Read(data)
	if err != nil {
		return color, err
	} else if n == 0 {
		return color, fmt.Errorf("missing data from %s", addr)
	}
	data = bytes.Trim(data, " \r\n")

	// parse data
	for i, p := range strings.Split(string(data), " ") {
		d, err := strconv.Atoi(p)
		if err != nil {
			return color, err
		}
		color[i] = Duty(d)
	}

	return color, nil
}

// GetPins returns a list with rgbw pins in use (-1 if not in use)
// TODO: update...
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

// TODO: update...
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

// TODO: update...
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
