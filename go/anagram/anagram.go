package anagram

import (
	"strings"
	"unicode"
)

type counts [26]int

//Detect which of the candidates is an Anagram
func Detect(subject string, candidates []string) (anagrams []string) {
	subjectCounts := count(subject)
	for _, candidate := range candidates {
		if !strings.EqualFold(subject, candidate) && subjectCounts == count(candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return
}

func count(s string) (c counts) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			c[unicode.ToLower(r)-'a']++
		}
	}
	return
}
