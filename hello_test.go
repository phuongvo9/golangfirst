package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Phuong")
		want := "Hello, Phuong"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' if no name is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})

}
