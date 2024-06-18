package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type languageBot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)

	fmt.Println(eb.getBye())
}

// func printGreeting(bot englishBot) {
// 	fmt.Println(bot.getGreeting())
// }

// func printGreeting(bot spanishBot) {
// 	fmt.Println(bot.getGreeting())
// }

func printGreeting(bot languageBot) {
	fmt.Println(bot.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (englishBot) getBye() string {
	return "By"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
