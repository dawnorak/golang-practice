package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter a word of your choice: ")
	var word string
	fmt.Scanln(&word)

	count := 0
	for _, char := range strings.ToLower(word) {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}

	fmt.Printf("The given word %s has %d letters.\n", word, count)
}