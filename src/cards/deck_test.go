package main

import "testing"

// comman to run test
// go test

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length to be 12 but got %v", len(d))
	}
}
