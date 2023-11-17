package main

import "fmt"

func main() {
	cards := deck{newCards(), "Three of Diamonds"}
	cards = append(cards, "Ace of Soldiers")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println((cards))
}
func newCards() string {
	return "Ace of Spades"
}
