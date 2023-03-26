package pico

import (
	"fmt"
	"net"
)

// GpPWM
type GpPWM struct {
	Nr   int `json:"nr"`   // Nr of gpio pin in use (gp0 - gp28)
	Duty int `json:"duty"` // Duty cycle (goes from 0-100)
}

// NewGpPWM
func NewGpPWM(nr int) *GpPWM {
	return &GpPWM{
		Nr: nr,
	}
}

// Device
type Device struct {
	ID   int       `json:"id"`   // ID is unique
	Addr string    `json:"addr"` // Addr contains the ip and port <ip>:<port>
	RGBW [4]*GpPWM `json:"rgbw"` // RGBW holds all pins in use
}

// NewDevice
func NewDevice(id int, addr string, rgbw [4]*GpPWM) *Device {
	return &Device{
		ID:   id,
		Addr: addr,
		RGBW: rgbw,
	}
}

// Handler for pico devices (and data)
type Handler struct {
	Devices []*Device `json:"devices"`
}

// NewPico
func NewHandler(devices ...*Device) *Handler {
	return &Handler{
		Devices: devices,
	}
}

// Scan for pico devices
func (h *Handler) Scan() (devices []*Device, err error) {
	// TODO: Scan the local network for pico devices
	// first get the current (local) ip address (default scan 192.168.178.0)
	localIP := localIP()
	if localIP == "" {
		return devices, fmt.Errorf("Local ip address not found!")
	}

	return devices, nil
}

func localIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String()
			}
		}
	}

	return ""
}
