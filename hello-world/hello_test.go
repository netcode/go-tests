package main

import "testing"

func TestHello(t *testing.T) {

	assertMsg := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertMsg(t, got, want)
	})

	t.Run("Say hello to someone", func(t *testing.T) {
		got := Hello("Eslam", "en")
		want := "Hello Eslam"
		assertMsg(t, got, want)
	})

	t.Run("Say hello world in spanish", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola World"
		assertMsg(t, got, want)
	})
	t.Run("Say hello to someone in spanish", func(t *testing.T) {
		got := Hello("Eslam", "es")
		want := "Hola Eslam"
		assertMsg(t, got, want)
	})

	t.Run("Say hello world in frensh", func(t *testing.T) {
		got := Hello("", "fr")
		want := "Bonjour World"
		assertMsg(t, got, want)
	})
	t.Run("Say hello to someone in frensh", func(t *testing.T) {
		got := Hello("Eslam", "fr")
		want := "Bonjour Eslam"
		assertMsg(t, got, want)
	})

}
