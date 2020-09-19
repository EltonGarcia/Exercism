package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ReplaceAll(strings.ToLower(phrase), ",", " ")
	reg := regexp.MustCompile(`([\w\s]|['][stre][\s])+`)
	freq := make(Frequency)
	for _, word := range reg.FindAllString(phrase, -1) {
		for _, s := range strings.Fields(word) {
			freq[s]++
		}
	}
	return freq
}
