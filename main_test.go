package main

import "testing"

func TestGreeting(t *testing.T) {
	if actual, expected := greeting(), "Hello, world!"; actual != expected {
		t.Fatalf("greeting() = %q, expected %q", actual, expected)
	}
}

func TestGreetingDoesNotReturnIncorrectMessage(t *testing.T) {
	if actual, unexpected := greeting(), "Goodbye, world!"; actual == unexpected {
		t.Fatalf("greeting() returned an incorrect message: %q", actual)
	}
}
