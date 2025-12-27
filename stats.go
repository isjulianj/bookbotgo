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

func countCharacters(text string) map[string]int {
	characterCounts := make(map[string]int)

	for _, c := range text {
		lower := strings.ToLower(string(c))
		if _, ok := characterCounts[lower]; !ok {
			characterCounts[lower] = 1
		} else {
			characterCounts[lower] += 1
		}

	}

	return characterCounts
}
