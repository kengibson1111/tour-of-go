package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordmap := make(map[string]int)

	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		count := wordmap[words[i]]
		wordmap[words[i]] = count+1
	}

	return wordmap
}

func main() {
	wc.Test(WordCount)
}

