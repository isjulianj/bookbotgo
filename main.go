package main

import (
	"fmt"
	"os"
)

func main() {
	path := "books/frankenstein.txt"

	text := get_book_text(path)

	fmt.Println(text)

}

func get_book_text(path string) (text string) {

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	text = string(file)
	return text
}
