package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	colors["red"] = "##ff00"
	// colors := map[string]string{
	// 	"red": "#ff000",
	// }

	// delete(colors, "red")

	printMap(colors)

	// fmt.Println((colors))
}

func printMap(colors map[string]string) {
	for color, hex := range colors {
		fmt.Println(color ,"-" ,hex)
	}
}
