package main

import (
	"bytes"
	"testing"
	"time"
)

type mockSleep struct {
	Calls int
}

func (ms *mockSleep) Sleep(d time.Duration) {
	ms.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := mockSleep{}
	Countdown(buffer, &sleeper)
	got := buffer.String()
	want := `5 
4 
3 
2 
1 
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
