package main

import "fmt"

func main() {

	// var card string = "this is card of spade"
	// card := "Five of Spade"
	// card = newCard()
	// fmt.Println(card)

	// printState()

	fmt.Println("slice example started")

	cards := deck{"Ace of Spades", "Five of Heart"}

	cards = append(cards, "King of Heart", newCard())

	fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Println(i, "index, card:", card)
	// }

	cards.printEach()

}

func newCard() string {
	return "Ace of Diamond"
}
