package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades' but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs' as last card but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("test_file.txt")

	d := newDeck()
	d.saveToFile("test_file.txt")

	loadedDeck := newDeckFromFile("test_file.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, go %v", len(loadedDeck))
	}
}
