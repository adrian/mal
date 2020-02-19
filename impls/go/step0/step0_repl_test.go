package main

import (
	"testing"
)

func TestRep(t *testing.T) {
	input := "dummy input"
	output := rep(input)
	if output != input {
		t.Errorf("expected %s but got %s", input, output)
	}
}