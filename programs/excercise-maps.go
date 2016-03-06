package main

import (
	//	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var word_count map[string]int = make(map[string]int)
	for _, word := range strings.Fields(s) {
		word_count[word] += 1
	}
	//	fmt.Println(s)
	//	fmt.Println(word_count)
	return word_count
}

func main() {
	wc.Test(WordCount)
}
