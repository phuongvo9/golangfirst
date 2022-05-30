package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Phuong")
		want := "Hello, Phuong"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello world' if no name is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
