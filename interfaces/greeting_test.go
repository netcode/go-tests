package interfaces

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Eslam")

	got := buffer.String()
	want := "Hello, Eslam"

	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}
