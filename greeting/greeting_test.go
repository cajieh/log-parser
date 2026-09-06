package greeting

import "testing"

func TestMessage(t *testing.T) {
	if actual, expected := Message(), "Hello, world!"; actual != expected {
		t.Fatalf("Message() = %q, expected %q", actual, expected)
	}
}

func TestMessageDoesNotReturnIncorrectMessage(t *testing.T) {
	if actual, unexpected := Message(), "Goodbye, world!"; actual == unexpected {
		t.Fatalf("Message() returned an incorrect message: %q", actual)
	}
}