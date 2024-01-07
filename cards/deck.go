package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
// A function to help us convert the deck type to byte slice
func (d deck) toString() string{
	return strings.Join([]string(d),",")
}
// Write File 
func (d deck) saveToFile() error{
	data := []byte(d.toString())
	return ioutil.WriteFile("test.csv",data,0666)
}
// Read File
func readFromFile(filename string) []string{
	// ReadFile(filename) will return 2 things : byte slice and error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	return strings.Split(string(bs)," ")

}