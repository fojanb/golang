package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
func (d deck) saveDeckToFile() error{
	data := []byte(d.toString())
	return ioutil.WriteFile("test.csv",data,0666)
}
// Read File
func readDeckFromFile(filename string) deck{
	// ReadFile(filename) will return 2 things : byte slice and error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs),","))
}
// func (d deck) shuffle() deck{
// 	for index := range d {
// 		r := rand.Intn(len(d))
// 		d[index] , d[r] = d[r] , d[index]
// 	}
// 	return d
// }
//A better approach to generate random number: seed-base
func (d deck) shuffle() deck{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newIndex := r.Intn(len(d))
		d[index] , d[newIndex] = d[newIndex] , d[index]
	}
	return d
}