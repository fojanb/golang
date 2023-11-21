package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
)

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
// print() is the receiver, reciever is a function or method whatever you like
// Note that print() and toString() function will be accessed by cards (or any variable of type deck)
// via . (dot)
// but the deal() function do not need . (dot)
func (d deck) print(){
	for index,card:=range d{
		fmt.Println(index,card)
	}
}
func deal(d deck, handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}
// ------------WriteFile------------
// toString() is a type convertor helper:
func (d deck) toString() string{
	return strings.Join([]string(d),",")
}
// WriteFile in Go:
func (d deck) saveToFile(filename string) error{
	return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}
// ------------ReadFile------------
// ReadFile in Go:
func readDeckFromFile(filename string) deck{
	// ReadFile() will return back 2 things: byte slice and error
	bs,err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	// note: []string(...) vs []string{...}, the first one is type conversion and the second one is
	// the slice of string initialization
	s := []string(strings.Split(string(bs),","))
	
	return deck(s)
}
// shuffling the slice of cards
func (d deck) shuffle(size int) deck{
	sliceOfCards := d[:size]
	for i:=range sliceOfCards{
		newi:=rand.Intn(len(sliceOfCards)-1)
		sliceOfCards[i],sliceOfCards[newi] = sliceOfCards[newi],sliceOfCards[i] 
	}
	return sliceOfCards
}