package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strings := strings.Fields(s)
	wcMap := make(map[string]int)

	for i := 0; i < len(strings); i++ {
		wcMap[strings[i]]++
	}

	return wcMap
}

func main() {
	wc.Test(WordCount)
}
