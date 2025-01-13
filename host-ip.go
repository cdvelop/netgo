package netgo

import (
	"errors"
	"fmt"
	"net"
)

// GetHostIP gets the host IP address
func GetHostIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("error getting network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		// Skip loopback and inactive interfaces
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			// Use only IPv4
			if ip.To4() != nil {
				return ip.String(), nil
			}
		}
	}

	return "", errors.New("could not get a valid IP address")
}
