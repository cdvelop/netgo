package netgo

import (
	"fmt"
	"net"
	"testing"
)

func TestAssignIPBasedOnCurrentMachine(t *testing.T) {
	// Get actual host IP
	hostIP, err := GetHostIP()
	if err != nil {
		t.Fatalf("Failed to get host IP: %v", err)
	}

	// Extract network portion of IP
	ip := net.ParseIP(hostIP).To4()
	if ip == nil {
		t.Fatalf("Invalid IPv4 address: %s", hostIP)
	}
	network := fmt.Sprintf("%d.%d.%d.0/24", ip[0], ip[1], ip[2])

	tests := []struct {
		name        string
		lastNumber  string
		wantIP      string
		wantNetmask string
		wantErr     bool
	}{
		{
			name:        "successful IP assignment",
			lastNumber:  "100",
			wantIP:      fmt.Sprintf("%d.%d.%d.100", ip[0], ip[1], ip[2]),
			wantNetmask: network,
			wantErr:     false,
		},
		{
			name:        "invalid last number",
			lastNumber:  "256", // Invalid IP octet
			wantIP:      "",
			wantNetmask: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotNetmask, err := AssignIPBasedOnCurrentMachine(tt.lastNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssignIPBasedOnCurrentMachine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIP != tt.wantIP {
				t.Errorf("AssignIPBasedOnCurrentMachine() gotIP = %v, want %v", gotIP, tt.wantIP)
			}
			if gotNetmask != tt.wantNetmask {
				t.Errorf("AssignIPBasedOnCurrentMachine() gotNetmask = %v, want %v", gotNetmask, tt.wantNetmask)
			}
		})
	}
}
