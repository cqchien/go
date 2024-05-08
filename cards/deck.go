package cards

import (
	"math/rand"
	"os"
	"strings"
	// "time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func (d deck) print() {
	for index, card := range d {
		println(index, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, d.toByteSlice(), 0666)
}

func newDeckFromFile(filename string) deck {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(fileContent), ", "))
}

func (d deck) shuffle() {
	// Gen a new seed that is different every time
	// source := rand.NewSource(time.Now().UnixNano())
	// newRand := rand.New(source)

	for index := range d {
		// Generate a random number between 0 and len(d) - 1 (not inclusive) -> output is the same every time because the seed is always the same
		newPosition := rand.Intn(len(d) - 1)
		// newPosition := newRand.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func newDeck() deck {
	cards := deck{}

	cardsSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardsSuites {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func getCardsWithNewType() {
	// Init a slice of strings
	cards := deck{"Ace of Spades", NewCard()}

	// Add an element to the slice
	cards = append(cards, "Six of Spades")

	// Iterate over the slice
	cards.print()
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
