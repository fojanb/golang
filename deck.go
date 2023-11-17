package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings
// next line means : deck ========= []string
type deck []string // deck is a type of slice of string(or array of strings)
func newDeck() deck{
	cards:=deck{}
	cardSuits:=[]string{"Spades","Hearts","Diamonds","Clubs"}
	cardValues:=[]string{"Ace","Two","Three","Four"}
	for _,suit:= range cardSuits{
		for _,value:=range cardValues{
			cards = append(cards,value + " of " + suit)
		}
	}
	return cards
}
// Any variable of type 'deck' now gets access to the 'print' method
// print() is the receiver
func (d deck) print(){
	for index,card:=range d{
		fmt.Println(index,card)
	}
}
func deal(d deck, handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}
