package isogram

import (
	"unicode"
)

//IsIsogram - Check if a word is a isogram
func IsIsogram(word string) bool {
	letters := make(map[rune]int)
	for _, char := range word {
		l := unicode.ToLower(char)
		count := letters[l]
		if count == 1 && unicode.IsLetter(l) {
			return false
		}
		letters[l] = count + 1
	}
	return true
}
