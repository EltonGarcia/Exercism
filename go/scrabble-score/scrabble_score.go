package scrabble

import (
	"unicode"
)

func letterValue(letter rune) int {
	switch unicode.ToUpper(letter) {
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 1
	}
}

//Score - computes the Scrabble score for a word.
func Score(word string) int {
	score := 0
	for _, w := range word {
		score += letterValue(w)
	}
	return score
}
