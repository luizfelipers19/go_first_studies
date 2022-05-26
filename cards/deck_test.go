package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52 cards, but got %d", len(d))
	}
}

func TestFirstAndLastCardsOfDeck(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' as the first card of the deck and got: %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected 'King of Diamonds' as the first card of the deck and got: %v", d[len(d)-1])
	}
}

func TestSaveToFileAndFileToDeck(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := deckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected a deck size of 52 cards and got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
