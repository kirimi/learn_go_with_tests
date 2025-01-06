package main

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("test search", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}

		got, err := dict.Search("test")
		want := "this is test"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("test search not found", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}

		_, err := dict.Search("keynotexist")

		assertError(t, err, ErrNotFound)
	})

	t.Run("test add", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Add("newkey", "new value")
		assertError(t, err, nil)

		val, err := dict.Search("newkey")
		assertError(t, err, nil)
		assertStrings(t, val, "new value")
	})

	t.Run("test add exist", func(t *testing.T) {
		dict := Dictionary{"test": "this is test"}

		err := dict.Add("test", "this is test")
		assertError(t, err, ErrKeyExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
