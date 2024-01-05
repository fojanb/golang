// OO Aproach vs GO Approach
package main
import "fmt"
func main(){
	// How to declare a new slice?
	cards := deck{"Ace of Spades",newCard()}
	// How to add new data to an existing slice?
	cards = append(cards,"Six of Diamonds")
	// How to loop through a slice in GO?
	for i,card := range cards {
		fmt.Println(i,card)
	}

}
func newCard() string {
	return "Five of Spades"
}