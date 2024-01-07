// Type Conversion in GO - We need this to use WriteFile function inside our application
package main
import "fmt"
func main(){
	// var greeting string= "Hello"
	// fmt.Println([]byte(greeting))
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//fmt.Println([]byte(cards.toString()))
	//cards.saveToFile()
	fmt.Println(readFromFile("test"))
}
