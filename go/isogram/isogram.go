package isogram

import (
	"unicode"
)

//IsIsogram - Check if a word is a isogram
func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, char := range word {
		l := unicode.ToLower(char)
		if unicode.IsLetter(l) {
			seen := letters[l]
			if seen {
				return false
			}
			letters[l] = true
		}
	}
	return true
}