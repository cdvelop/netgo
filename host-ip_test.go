package netgo

import (
	"net"
	"testing"
)

func TestGetHostIP(t *testing.T) {
	ip, err := GetHostIP()
	if err != nil {
		t.Fatalf("GetHostIP() failed: %v", err)
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		t.Errorf("GetHostIP() returned invalid IP address: %s", ip)
	}

	if parsedIP.IsLoopback() {
		t.Errorf("GetHostIP() returned loopback address: %s", ip)
	}
}
