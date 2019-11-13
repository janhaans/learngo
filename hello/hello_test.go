package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	checkResult := func(t *testing.T, got string, expected string) {
		t.Helper()

		if got != expected {
			t.Errorf("Result is not OK, expected %q but got %q", expected, got)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Jan", "")
		expected := "Hello, Jan"

		checkResult(t, got, expected)
	})

	t.Run("Say hello to world when no name is given", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, world"

		checkResult(t, got, expected)
	})

	t.Run("Say hello to people in Dutch", func(t *testing.T) {
		got := Hello("Kees", "dutch")
		expected := "Hallo, Kees"

		checkResult(t, got, expected)
	})

}
