package main

import "fmt"

// Struct is a key value pair like object in javascript
type contactInfo struct{
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
	//contactInfo contactInfo
}

func main() {
// #1 struct declaration:
	// alex := person{"Alex", "Anderson"}
	// fmt.Println(alex)
// #2 struct declaration:
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
// #3 struct declaration:	
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v",alex)
	alex := person{
		firstName: "Alex",
		lastName: "Party",
		contactInfo: contactInfo{
			email:"alex@gmail.com",
			zipCode:1000,
		},
	}
	//fmt.Printf("%+v",alex)
	//fmt.Print(alex)
	// now since person is a custom type, so we can write a receiver for it :
	alex.print()
}
func (p person) print(){
	fmt.Printf("%+v",p)
}
