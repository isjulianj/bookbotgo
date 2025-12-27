package main

import (
	"fmt"
)

func main() {
	path := "books/frankenstein.txt"

	text := getBookText(path)

	fmt.Println(countWords(text))
	mappy := countCharacters(text)

	for key, value := range mappy {
		fmt.Printf("%s : %d\n", key, value)
	}
}
