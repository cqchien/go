package main

import "fmt"

func evenOrOdd() {
	arrNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range arrNumber {
		if isEvenNumber(number) {
			fmt.Println(number, "is even number")
		} else {
			fmt.Println(number, "is odd number")
		}
	}
}

func isEvenNumber(number int) bool {
	return number%2 == 0
}
