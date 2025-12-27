package main

import (
	"os"
	"strings"
)

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
