package summultiples

import (
	"fmt"
	"math"
)

func SumMultiples(limit int, values ...int) int {
	limit--
	var total int
	for _, v := range values {
		total += sumMultiples(limit, v)
	}
	fmt.Printf("Total: %d\n", total)
	if len(values) <= 1 {
		return total
	}

	cd := findGcd(values)
	if cd == 1 {
		// hasNonPrimeFactor := false
		// for i := 0; i < len(values); i++ {
		// 	for j := i + 1; j < len(values); j++ {
		// 		a := values[i]
		// 		b := values[j]
		// 		cd = gcd(a, b)
		// 		if cd == 1 {
		// 			total -= sumMultiples(limit, a*b)
		// 		} else {
		// 			// hasNonPrimeFactor = true
		// 			// total -= sumMultiples(limit, lcm(a, b))
		// 			// fmt.Printf("GCD(%d,%d) = %d\n", a, b, cd)
		// 			// fmt.Printf("Sum of GCD(%d,%d) = %d\n", a, b, sumMultiples(limit, cd))
		// 			// fmt.Printf("LCM(%d,%d) = %d\n", a, b, lcm(a, b))
		// 			// fmt.Printf("Sum of LCM(%d,%d) = %d\n", a, b, sumMultiples(limit, lcm(a, b)))
		// 			// fmt.Printf("Total: %d, Diff: %d\n", total, total-4419)
		// 		}
		// 	}
		// }
		totalPrimes := totalOfCombinations(limit, values...)
		fmt.Printf("Total primes: %d\n", totalPrimes)
		total += totalPrimes
		fmt.Printf("Total: %d\n", total)

		// if hasNonPrimeFactor {
		// 	cm := findLcm(values)
		// 	total += sumMultiples(limit, cm)
		// }
	} else {
		cm := findLcm(values)
		total -= sumMultiples(limit, cm)
	}

	return total
}

func totalOfCombinations(limit int, values ...int) int {
	total := 0
	for i := 2; i <= len(values); i++ {
		combinations := createCombinations(i, values...)
		for _, comb := range combinations {
			x := findLcm(comb)
			value := sumMultiples(limit, x)
			signal := int(math.Pow(-1, float64(i-1)))
			fmt.Printf("%v: LCM:%d, Total:%d\n", comb, x, signal*value)
			total += signal * value
		}
	}
	return total
}

func createCombinations(size int, values ...int) [][]int {
	n := len(values)
	count := factorial(n) / (factorial(size) * factorial(n-size))
	combinations := make([][]int, count)
	pointers := make([]int, size)
	for i := 0; i < size; i++ {
		pointers[i] = i
	}
	moving := size - 1
	for i := 0; i < count; i++ {
		item := make([]int, size)
		for i := 0; i < size; i++ {
			item[i] = values[pointers[i]]
		}
		combinations[i] = item
		if pointers[moving] == moving+n-size {
			for moving--; moving >= 0; moving-- {
				if pointers[moving] == moving+n-size {
					continue
				}
				pointers[moving]++
				next := moving + 1
				for j := pointers[moving] + 1; j < n && next < size; j++ {
					pointers[next] = j
					if pointers[next] != n-size+next {
						moving = next
					}
					next++
				}
				break
			}
		} else {
			pointers[moving]++
		}
	}
	return combinations
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func sumMultiples(limit int, number int) int {
	if number == 0 {
		return 0
	}
	multiples := (limit - (limit % number)) / number
	return number * ((1 + multiples) * multiples / 2)
}

func findLcm(values []int) int {
	result := 1
	for i := 0; i < len(values); i++ {
		result = lcm(result, values[i])
	}
	return result
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func findGcd(values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		result = gcd(result, values[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
