package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	ans := get_nearest_5min(1)
	if ans != 0 {
		t.Fatalf("failed")
	}
}
