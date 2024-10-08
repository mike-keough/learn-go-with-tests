package main

import "testing"

func TestHello(t *testing.T) {
  
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Mike", "")
		want := "Hello, Mike"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello in spanish", func(t *testing.T) {
		got := Hello("Mike", "Spanish")
		want := "Hola, Mike"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello in french", func(t *testing.T) {
		got := Hello("Mike", "French")
		want := "Bonjour, Mike"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
