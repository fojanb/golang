package main
import "fmt"

//Create a type of 'deck'
//which is a slice of string
type deck []string

// Think of print() as a method in a class
func (d deck) print(){
	for i,card := range d {
		fmt.Println(i,card)
	}
}
func newDeck() deck{
	cards := deck{}
	suits := []string{"Spades","Diamonds","King"}
	vals := []string{"Ace","Two","Three"}
	for _,suit := range suits{
		for _,val := range vals{
			card := val +" of "+suit
			cards = append(cards,card)
		}
	}
	return cards
}
func (d deck) deal(handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
	// hand,remains := d[:handSize],d[handSize:]
}
//Or:
/*
func deal(d deck , handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}

*/
