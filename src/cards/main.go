package main

import "fmt"

func main() {

	// var card string = "this is card of spade"
	// card := "Five of Spade"
	// card = newCard()
	// fmt.Println(card)

	// printState()

	fmt.Println("slice example started")

	// cards := deck{"Ace of Spades", "Five of Heart"}

	// cards = append(cards, "King of Heart", newCard())

	// fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, "index, card:", card)
	// }

	// fmt.Println("Type conversion example")
	// greeting := "Hi there!"

	// fmt.Println([]byte(greeting))

	cards := newDeck()

	fmt.Println("priting the original deck")
	cards.printEach()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println("priting the deal deck")
	// hand.printEach()

	// fmt.Println("priting the leftover deck")
	// remainingDeck.printEach()

	// fmt.Println("printing cardds into string")
	// cardString := cards.toString()

	// fmt.Println(cardString)

	// cards.saveToFile("my_cards")

	// newCards := newDeckFromFile("my_cards")

	// fmt.Println("+++++ New cards from file ++++++")
	// newCards.printEach()

	cards.suffle()

	fmt.Println("suffle cards")
	cards.suffle()
	cards.printEach()
}

// func newCard() string {
// 	return "Ace of Diamond"
// }
