package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "books/frankenstein.txt"

	text := getBookText(path)

	fmt.Println(countWords(text))

}

func getBookText(path string) (text string) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text = string(file)
	return text
}

func countWords(text string) int {
	ss := strings.Fields(text)

	return len(ss)
}
