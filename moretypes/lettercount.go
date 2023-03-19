package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a word of your choice: ")
	var word string
	fmt.Scanln(&word)

	count := 0
	for _, char := range (word) {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}

	fmt.Printf("The given word %s has %d letters.\n", word, count)
}