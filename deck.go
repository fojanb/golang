package main

import "fmt"

// Create a new type of 'deck', which is a slice of strings
// next line means : deck ========= []string
type deck []string // deck is a type of slice of string(or array of strings)
// Any variable of type 'deck' now gets access to the 'print' method
func (d deck) print(){
	for index,card:=range d{
		fmt.Println(index,card)
	}
}
