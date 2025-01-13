package netgo

import (
	"testing"
)

func TestPing(t *testing.T) {
	err := Ping("8.8.8.8")
	if err != nil {
		t.Errorf("Ping() error = %v, want no error", err)
	}
}
