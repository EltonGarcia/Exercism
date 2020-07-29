package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	size := len(rhyme)
	rhymes := make([]string, size)
	for i := 0; i < len(rhyme)-1; i++ {
		rhymes[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	if size > 0 {
		rhymes[size-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	}
	return rhymes
}
