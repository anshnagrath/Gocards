package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}
	if d[0] != "Ace of spade" {
		t.Errorf("Expected first card of Ace of spade but got %v", d[0])
	}
	if d[len(d)-1] != "four of club" {
		t.Errorf("Expected last card four of club got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(loadedDeck))
	}
	defer os.Remove("_decktesting")

}
