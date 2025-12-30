package main

import (
	"fmt"
)

func main() {
	path := "books/frankenstein.txt"

	text := getBookText(path)

	charMap := countCharacters(text)
	sortedInts := getSortedCharacters(charMap)

	fmt.Println("============ BOOKBOT ============")
	fmt.Println("Analyzing book found at books/frankenstein.txt...")
	fmt.Println("----------- Word Count ----------")
	fmt.Printf("Found %d total words\n", countWords(text))
	fmt.Println("--------- Character Count -------")
	for _, c := range sortedInts {
		fmt.Printf("%s : %d\n", c.character, c.count)
	}
	fmt.Println("============= END ===============")
}
