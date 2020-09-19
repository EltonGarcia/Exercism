package summultiples

import (
	"fmt"
	"testing"
)

func TestSumMultiples(t *testing.T) {
	for _, test := range varTests {
		s := SumMultiples(test.limit, test.divisors...)
		if s != test.sum {
			t.Fatalf("Sum of multiples of %v to %d returned %d, want %d.",
				test.divisors, test.limit, s, test.sum)
		}
	}
}

func BenchmarkSumMultiples(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiples(test.limit, test.divisors...)
		}
	}
}

func TestCreateCombinations(t *testing.T) {
	values := []int{2, 3, 9, 11, 17, 21}
	combinations := createCombinations(1, values...)
	for i, comb := range combinations {
		fmt.Printf("%d: %v\n", i+1, comb)
	}
}
