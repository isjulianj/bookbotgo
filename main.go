package main

import (
	"fmt"
)

func main() {
	path := "books/frankenstein.txt"

	text := getBookText(path)

	fmt.Println(countWords(text))

}
