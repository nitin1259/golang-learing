package main

import "fmt"

// create a new type of deck
// which is slice of string.

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Heart", "Diamond", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) printEach() {
	for i, card := range d {
		fmt.Println("index: ", i, "card: ", card)
	}
}
