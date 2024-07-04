package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	msg := ""
	if msg != "" {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
