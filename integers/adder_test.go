package main

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf(" Expected %q got %q", expected, sum)
	}
}
