package etl

import (
	"strings"
)

//Legacy - The legacy format
type Legacy map[int][]string

//Shiny - The new shiny format
type Shiny map[string]int

//Transform - transform the legacy format to the new shiny format
func Transform(source Legacy) Shiny {
	var result Shiny = make(map[string]int)
	for score, letters := range source {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}
	return result
}
