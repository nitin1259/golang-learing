package main

import "fmt"

// create a new type of deck
// which is slice of string.

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Heart", "Diamond", "Clubs"}

	// cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cardValues := []string{"Ace", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) printEach() {
	for i, card := range d {
		fmt.Println("index: ", i, "card: ", card)
	}
}
