package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to someone specific", func(t *testing.T) {
		got := Hello("Vini")
		want := "Hello, Vini"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
