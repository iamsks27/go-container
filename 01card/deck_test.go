package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16 cards, got %d", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Removing file if already exists
	os.Remove("_decktesting")

	// Creating a new deck & saving it in a file
	d := newDeck()
	d.saveToFile("_decktesting")

	// Creating a new deck by reading the connect of a file
	ld := newDeckFromFile("_decktesting")

	// Running tests
	if len(ld) != 16 {
		t.Errorf("Expecting 16 cards, got %d", len(ld))
	}

	// Removing the test file
	os.Remove("_decktesting")
}
