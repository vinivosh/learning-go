package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("say hello to someone specific", func(t *testing.T) {
		got := Hello("Vini", "")
		want := "Hello, Vini"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to someone specific --- in Spanish", func(t *testing.T) {
		got := Hello("Pablo Escobar", "es")
		want := "Hola, Pablo Escobar"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to someone specific --- in French", func(t *testing.T) {
		got := Hello("Proudhon", "fr")
		want := "Bonjour, Proudhon"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to someone specific --- in Portuguese", func(t *testing.T) {
		got := Hello("Epaminondas", "pt")
		want := "Olá, Epaminondas"

		assertCorrectMessage(t, got, want)
	})
}
