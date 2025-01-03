package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		name := "Kirill"
		got := Hello(name, "en")
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when no name is given", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In russian", func(t *testing.T) {
		got := Hello("Кирилл", "ru")
		want := "Привет, Кирилл"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when no name is given, in russian", func(t *testing.T) {
		got := Hello("", "ru")
		want := "Привет, Мир"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q - Want %q", got, want)
	}
}
