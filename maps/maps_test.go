package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is testing", "val": "this is val"}

	t.Run("word found", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		wants := "this is testing"

		assertString(t, got, wants)
	})

	t.Run("word notfound", func(t *testing.T) {
		_, err := dictionary.Search("noooo")
		if err == nil {
			t.Errorf("expected error, didnt find any")
		}
		if err != ErrWordNotFound {
			t.Errorf("Got different error than expected: error: %s", err)
		}
	})

}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{"test": "this is testing"}
	t.Run("adding new word", func(t *testing.T) {
		def := "Just a value"
		dictionary.Add("val", def)
		got, _ := dictionary.Search("val")
		assertString(t, got, def)
	})

	t.Run("adding existing word", func(t *testing.T) {
		def := "Just a value"
		err := dictionary.Add("test", def)
		if err != ErrWordExists {
			t.Errorf("Got different error than expected: error: %s", err)
		}
	})

}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is testing"}

	t.Run("update existing word", func(t *testing.T) {
		def := "this is testing too."
		dictionary.Update("test", def)
		got, _ := dictionary.Search("test")
		assertString(t, got, def)
	})

	t.Run("update not found word", func(t *testing.T) {
		def := "Just a value"
		err := dictionary.Update("val", def)
		if err != ErrWordNotFound {
			t.Errorf("Got different error than expected: error: %s", err)
		}
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "this is just a test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Word is still exists %s", word)
	}
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
