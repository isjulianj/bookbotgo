package main

import (
	"fmt"
	"os"
)

func main() {
	path := "books/frankenstein.txt"

	text := getBookText(path)

	fmt.Println(text)

}

func getBookText(path string) (text string) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text = string(file)
	return text
}
