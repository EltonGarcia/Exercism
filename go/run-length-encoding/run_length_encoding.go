package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	var result string
	lastChange := 0
	for i := 1; i <= len(input); i++ {
		if i == len(input) || input[lastChange] != input[i] {
			count := i - lastChange
			occurences := ""
			if count > 1 {
				occurences = strconv.Itoa(count)
			}
			result += fmt.Sprintf("%s%s", occurences, string(input[lastChange]))
			lastChange = i
		}
	}
	return result
}

func RunLengthDecode(input string) string {
	var result string
	ocurrences := 0
	for _, c := range input {
		if unicode.IsNumber(c) {
			ocurrences = (ocurrences * 10) + int(c-'0')
		} else {
			if ocurrences == 0 {
				ocurrences = 1
			}
			result += strings.Repeat(string(c), ocurrences)
			ocurrences = 0
		}
	}
	return result
}
