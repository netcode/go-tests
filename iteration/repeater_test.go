package main

import "testing"

func TestIteration(t *testing.T) {
	result := Repeater("a", 5)
	expected := "aaaaa"
	if result != expected {
		t.Errorf("Got %q expected %q", result, expected)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a", 5)
	}
}
