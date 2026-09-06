package main

import "testing"

func TestGreeting(t *testing.T) {
	if got, want := greeting(), "Hello, world!"; got != want {
		t.Fatalf("greeting() = %q, want %q", got, want)
	}
}