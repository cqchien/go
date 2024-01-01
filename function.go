package main

import "fmt"

func getCard() {
	card := newCard();

	fmt.Println(`This is a card:`, card)
}

func newCard() string {
	// Init a slice of strings
	card := "Ace of Spades"

	// Return the slice
	return card
}