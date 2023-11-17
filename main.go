package main
func main() {
	cards := deck{newCards(), "Three of Diamonds"}
	cards = append(cards, "Ace of Soldiers")
	makan := deck{"fname:Makan","lname:Fard"}
	cards.print()
	makan.print()
}
func newCards() string {
	return "Ace of Spades"
}
