package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 32 {
		t.Errorf("Expected deck length of 32, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected element is Ace of Clubs but found %v", d[0])
	}

	if d[len(d)-1] != "Eight of Spades" {
		t.Errorf("Expected element is Ace of Clubs but found %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fn := "_decktesting"
	os.Remove(fn)
	deck := newDeck()
	deck.saveToFile(fn)
	loadedDeck := newDeckFromFile(fn)
	if len(loadedDeck) != 32 {
		t.Errorf("Expected 32 cards and deck got %v", len(loadedDeck))
	}
	os.Remove(fn)

}
