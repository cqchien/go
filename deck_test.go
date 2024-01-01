package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Check if the deck has the correct number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Check if the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	// Check if the last card is Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Delete any existing file
	os.Remove("_decktesting")

	// Create a new deck
	d := newDeck()

	// Save the deck to file
	d.saveToFile("_decktesting")

	// Load the deck from file
	loadDeck := newDeckFromFile("_decktesting")

	// Check if the deck has the correct number of cards
	if len(loadDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadDeck))
	}

	// Delete the file
	os.Remove("_decktesting")
}
