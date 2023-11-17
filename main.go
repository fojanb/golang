package main
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
	hand.print()
	remaining.print()
}
// func newCards() string {
// 	return "Ace of Spades"
// }

// Important example:
/*
package main
 
import "fmt"
 
func main() {
   c := color("Red")
 
   fmt.Println(c.describe("is an awesome color"))
}
 
type color string
 
func (c color) describe(description string) (string) {
   return string(c) + " " + description
}

*/
