package cards

import "fmt"

func GetCard() {
	card := NewCard();

	fmt.Println(`This is a card:`, card)
}

func NewCard() string {
	// Init a slice of strings
	card := "Ace of Spades"

	// Return the slice
	return card
}