package pico

import (
	"bytes"
	"fmt"
	"net"
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

func RunCommand(addr string, command string, readSize int) (data []byte, err error) {
	// dial to picow device
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return data, &DialError{err}
	}
	defer conn.Close()

	// send command
	n, err := conn.Write([]byte(command))
	if err != nil {
		return data, err
	} else if n == 0 {
		return data, fmt.Errorf("no data written to %s", addr)
	}

	// read response into data
	if readSize <= 0 {
		// skip read data
		return data, nil
	}

	data = make([]byte, readSize)
	n, err = conn.Read(data)
	if err != nil {
		return data, err
	} else if n == 0 {
		return data, fmt.Errorf("missing data from %s", addr)
	}
	data = bytes.Trim(data, " \r\n")

	return data, nil
}

// GetColor from picow device
func GetColor(addr string) (color [4]Duty, err error) {
	data, err := RunCommand(addr, TCPGetColorCommand, TCPGetColorReadSize)
	if err != nil {
		return color, err
	}

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

func SetColor(addr string, color [4]Duty) error {
	command := fmt.Sprintf(TCPSetColorCommand,
		color[0], color[1], color[2], color[3])
	_, err := RunCommand(addr, command, TCPSetColorReadSize)
	return err
}

// GetGpPins (GPIO) pins from picow device in use for rgbw
func GetGpPins(addr string) (gps [4]GpPin, err error) {
	data, err := RunCommand(addr, TCPGetGpPinsCommand, TCPGetGpPinsReadSize)
	if err != nil {
		return gps, err
	}

	// parse data
	for i, p := range strings.Split(string(data), " ") {
		d, err := strconv.Atoi(p)
		if err != nil {
			return gps, err
		}
		gps[i] = GpPin(d)
	}

	return
}

func SetGpPins(addr string, gps [4]GpPin) error {
	command := fmt.Sprintf(TCPSetGpPinsCommand,
		gps[0], gps[1], gps[2], gps[3])

	_, err := RunCommand(addr, command, TCPSetGpPinsReadSize)
	return err
}
