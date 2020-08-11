package luhn

import (
	"strings"
	"unicode"
)

//Valid - Check if a credit card number is valid
func Valid(card string) bool {
	card = strings.ReplaceAll(card, " ", "")
	if len(card) < 2 {
		return false
	}
	oddPos := len(card)%2 == 0

	var sum int
	for _, char := range card {
		if !unicode.IsNumber(char) {
			return false
		}
		digit := int(char - '0')
		if oddPos {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		oddPos = !oddPos
	}
	return sum%10 == 0
}
