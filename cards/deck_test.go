package main

import (
	"testing"
	"os"
)

//Write some codes to evalute the logics inside the deck.go file
func TestNewDeck(t *testing.T){
	d := newDeck()
	if len(d) != 9 {
		t.Errorf("We expected to get the a deck of length 9 but we got:%v",len(d))
	}
	if d[len(d)-1] != "Three of King" || d[0] != "Ace of Spades" {
		t.Errorf("The last or first card in the deck is not the expected card")
	}
}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T){
	os.Remove("_decktesting.txt")
	d := newDeck()
	d.saveDeckToFile("_decktesting.txt")
	loadedDeck := readDeckFromFile("_decktesting.txt")
	if len(loadedDeck) != 9 {
		t.Errorf("We expected a deck of length 15 but got %v",len(loadedDeck))
	}
	os.Remove("_decktesting.txt")
}