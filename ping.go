package netgo

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func Ping(ip string) error {

	// Ping checks connectivity with the VM
	if ip == "" {
		return errors.New("VM has no assigned IP")
	}

	// Execute ping (Windows)
	cmd := exec.Command("ping", "-n", "1", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing ping: %v", err)
	}

	// Check if there was a response
	if !strings.Contains(string(output), "Reply from") {
		return fmt.Errorf("no response from %s", ip)
	}

	return nil
}
