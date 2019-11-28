package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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

func (d deck) toString() string {

	// type conversion of deck type into slice of string - []string(d)
	cards := strings.Join([]string(d), ",")
	return cards
}

func (d deck) saveToFile(fileName string) error {
	// func WriteFile(filename string, data []byte, perm os.FileMode) error
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {

	// func ReadFile(filename string) ([]byte, error)
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		// option 1- log the error and return a new fresh deck
		// option 2 - log the error and exit the program.
		os.Exit(1)
	}
	//type conversion of byte slice
	cardString := string(bs) // cardString became Ace of Spades, Two of Spades, Three of Spades

	cards := strings.Split(cardString, ",")

	return deck(cards)

}

func (d deck) suffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
