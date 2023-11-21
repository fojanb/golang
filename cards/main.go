package main


func main() {
	// cards := deck{newCards(), "Three of Diamonds"}
	// cards = append(cards, "Ace of Soldiers")
	// makan := deck{"fname:Makan","lname:Fard"}
	// cards.print()
	// makan.print()
	// newDeck()
	// cards := newDeck()
	// cards.print()
	// fmt.Println("CARDS",cards)
	// hand,remaining := deal(cards,2)
	// hand.print()
	// remaining.print()
	//fmt.Println("Byte slice for Fojan : ",[]byte("Fojan"))
	//fmt.Println(cards.toString())
	//fmt.Println([]byte(cards.toString()))
	// cards.saveToFile("Fojan_Golang.txt")
	// myFile := readDeckFromFile("Fojan_Golang.txt")
	// fmt.Println("My file is:")
	// myFile.print()
	cards := newDeck()
	shuffledCards := cards.shuffle(6)
	shuffledCards.print()

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
