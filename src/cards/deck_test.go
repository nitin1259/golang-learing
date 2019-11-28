package main

import "testing"

// comman to run test
// go test

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length to be 12, but got %v", len(d))
	} else if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	} else if d[3] != "Ace of Heart" {
		t.Errorf("Expected first card to be 'Ace of Heart', but got %v", d[3])
	} else if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected first card to be 'Three of Clubs', but got %v", d[len(d)-1])
	}
}
