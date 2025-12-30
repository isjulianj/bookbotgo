package main

import (
	"os"
	"sort"
	"strings"
	"unicode"
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

func countCharacters(text string) (characterCounts map[string]int) {
	characterCounts = make(map[string]int)

	for _, c := range text {
		lower := strings.ToLower(string(c))
		if _, ok := characterCounts[lower]; !ok {
			characterCounts[lower] = 1
		} else {
			characterCounts[lower] += 1
		}
	}
	return
}

type characterCount struct {
	character string
	count     int
}

func getSortedCharacters(characterCountMap map[string]int) []characterCount {
	list := make([]characterCount, 0, len(characterCountMap))

	for key, value := range characterCountMap {
		if isLetter(key) {
			list = append(list, characterCount{character: key, count: value})
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})

	return list
}

func isLetter(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}
