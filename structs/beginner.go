package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	// Print struct with field names and values
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@hello.com",
			zipCode: 12345,
		},
	}

	fmt.Println(&alex)
	// alex.print();

	// alexPointer := &alex
	// alexPointer.updateFirstName("Chien")

	// Shortcut to update the value of a struct
	// alex.updateFirstName("Chien");
	// alex.print();
}
