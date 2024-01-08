package main
// import "fmt"

func main(){
	jim := person {
		firstName:"Jim", 
		lastName:"Anderson", 
		contactInfo: contactInfo {
			email:"fojan@gmail.com", 
			zip:123,
		},
	}
	jim.updateFirstName("Makan")
	jim.printPerson()
}