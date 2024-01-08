package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) printPerson() {
	fmt.Printf("%+v",p)

}
func (p *person) updateFirstName(newName string) {
	(*p).firstName = newName
}