package romannumerals

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

var letters = map[int]rune{
	1:    'I',
	5:    'V',
	10:   'X',
	50:   'L',
	100:  'C',
	500:  'D',
	1000: 'M',
}

//ToRomanNumeral - Converts arabic numbers to Roman Numerals
func ToRomanNumeral(number int) (string, error) {
	if number < 1 || number > 3000 {
		return "", errors.New("Invalid input")
	}
	var result strings.Builder
	digits := strconv.Itoa(number)
	length := len(digits)
	for i, c := range digits {
		n := int(c - '0')
		key := int(math.Pow10(length - i - 1))
		if n == 9 || n == 4 {
			result.WriteRune(letters[key])
			result.WriteRune(letters[key*(n+1)])
		} else {
			for n > 0 {
				if n >= 5 {
					result.WriteRune(letters[key*5])
					n -= 5
				} else {
					result.WriteRune(letters[key])
					n--
				}
			}
		}
	}
	return result.String(), nil
}
