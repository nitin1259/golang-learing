package main

import "fmt"

// create a new type of deck
// which is slice of string.

type deck []string

func (d deck) printEach() {
	for i, card := range d {
		fmt.Println("index: ", i, "card: ", card)
	}
}
