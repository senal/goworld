package main

import "testing"

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
