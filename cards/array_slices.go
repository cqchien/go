package cards

func GetCards() {
	// Init a slice of strings
	cards := []string{"Ace of Spades", NewCard()}

	// Add an element to the slice
	cards = append(cards, "Six of Spades")

	// Iterate over the slice
	for index, card := range cards {
		println(index, card)
	}
}