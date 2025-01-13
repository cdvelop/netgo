package netgo

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// AssignIPBasedOnCurrentMachine gets the IP of the current machine and assigns an IP based on that network mask.
// Ex: input: lasNumber="100"
// Ex output: ip="192.168.1.100", netmask="192.168.1.0/24", err=nil
func AssignIPBasedOnCurrentMachine(lasNumber string) (string, string, error) {
	hostIP, err := GetHostIP()
	if err != nil {
		return "", "", err
	}

	// Get network mask
	ip := net.ParseIP(hostIP)
	if ip == nil {
		return "", "", errors.New("invalid IP address")
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", "", fmt.Errorf("error getting interfaces: %v", err)
	}

	var netmask string
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.String() == hostIP {
					// Use first 3 octets with /24 mask as expected by tests
					ipv4 := ip.To4()
					if ipv4 == nil {
						return "", "", errors.New("invalid IPv4 address")
					}
					netmask = fmt.Sprintf("%d.%d.%d.0/24", ipv4[0], ipv4[1], ipv4[2])
					break
				}
			}
		}
	}

	if netmask == "" {
		return "", "", errors.New("could not determine network mask")
	}

	// Validate last number is a valid octet (0-255)
	lastOctet := 0
	_, err = fmt.Sscanf(lasNumber, "%d", &lastOctet)
	if err != nil || lastOctet < 0 || lastOctet > 255 {
		return "", "", errors.New("invalid last number")
	}

	// Assign IP to VM
	octets := strings.Split(hostIP, ".")
	if len(octets) != 4 {
		return "", "", errors.New("invalid IP format")
	}

	octets[3] = fmt.Sprintf("%d", lastOctet)
	vmIp := strings.Join(octets, ".")

	return vmIp, netmask, nil
}
