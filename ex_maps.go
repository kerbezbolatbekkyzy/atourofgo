package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Split the string into words
	words := strings.Fields(s)

	// Create a map to store word counts
	counts := make(map[string]int)

	// Count the occurrence of each word
	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
