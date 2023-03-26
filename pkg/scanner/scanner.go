package scanner

import (
	"fmt"
	"net"
)

// GetLocalIP address
func GetLocalIP() (string, error) {
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.To4().String(), nil
			}
		}
	}

	return "", fmt.Errorf("Local ip address not found!")
}
