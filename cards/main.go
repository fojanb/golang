// Shuffle function for deck
package main
// import "fmt"
func main(){
	cards := newDeck()
	cards.saveDeckToFile("test.txt")
}
