package maps

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	var word_count map[string]int = make(map[string]int)
	for _, word := range words {
		count := word_count[word]
		word_count[word] = count + 1
	}
	return word_count
}

func main() {
	wc.Test(WordCount)
}
