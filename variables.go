package main

import "fmt"

func getVariables() {
	// Init a variable with a type
	// var card string = "Ace of Spades"

	// Init a variable without a type (Go will infer the type)
	card := "Ace of Spades"

	// Reassign the variable
	card = "Five of Diamonds"

	// Assign a new variable first and then reassign it
	var name string

	name = "John Doe"
	fmt.Println(name)

	// Print the variable
	fmt.Println(card)
}
