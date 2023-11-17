package main

import "fmt"

func main() {
	// cards := deck{newCards(), "Three of Diamonds"}
	// cards = append(cards, "Ace of Soldiers")
	// makan := deck{"fname:Makan","lname:Fard"}
	// cards.print()
	// makan.print()
	// newDeck()
	cards := newDeck()
	cards.print()
	hand,remaining := deal(cards,2)
	fmt.Println("hand: ",hand,"remaining: ",remaining)
}
// func newCards() string {
// 	return "Ace of Spades"
// }
