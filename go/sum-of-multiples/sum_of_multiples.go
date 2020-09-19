package summultiples

import (
	"math"
)

//SumMultiples - Given a number, find the sum of all the unique
//multiples of particular numbers up to but not including that number.
func SumMultiples(limit int, values ...int) int {
	//Since the last number is not included
	//exclude the last number by subtracting one from the limit
	limit--
	//returns zero if no values are provided
	if len(values) < 1 {
		return 0
	}
	return totalOfCombinations(limit, values...)
}

//totalOfCombinations - calculates the sum of the combinations
//using the inclusion-exclusion principle:
//|A⋃B⋃C| = |A| + |B| + |C| - |A⋂B| - |A⋂C| - |B⋂C| + |A⋂B⋂C|
//i.e.:
//∑(-1)^n-1 * ∑|Ai⋂Ak|
func totalOfCombinations(limit int, values ...int) int {
	total := 0
	for i := 1; i <= len(values); i++ {
		combinations := createCombinations(i, values...)
		signal := int(math.Pow(-1, float64(i-1)))
		for _, comb := range combinations {
			combLcm := findLcm(comb)
			value := sumMultiples(limit, combLcm)
			total += signal * value
		}
	}
	return total
}

//createCombinations - Given the size and values creates all possible
//combinations of element the 'size' size
func createCombinations(size int, values ...int) [][]int {
	n := len(values)
	//calculates the number of possible combinations: n! / (s! * (n-s)!)
	count := factorial(n) / (factorial(size) * factorial(n-size))
	//array to keep all combinations
	combinations := make([][]int, count)
	//create pointer to keep track of the number inserted in a combination
	pointers := make([]int, size)
	//initialize the pointer in the first 'size' positions
	for i := 0; i < size; i++ {
		pointers[i] = i
	}
	//the last pointer starts to move to create different combinations
	moving := size - 1

	for i := 0; i < count; i++ {
		//creates the combination item
		item := make([]int, size)
		//assign the combinations values from the pointer
		for i := 0; i < size; i++ {
			item[i] = values[pointers[i]]
		}
		//assign the combination item to combinations list
		combinations[i] = item
		//if the current pointer moving has reach it's max position
		if pointers[moving] == moving+n-size {
			for {
				//set the moving pointer as the one before the current
				moving--
				//stop if it's not a valid pointer
				if moving < 0 {
					break
				}
				//if the current pointer is also on it's max position
				//continue with the pointer before the current
				if pointers[moving] == moving+n-size {
					continue
				}
				//advance the current pointer moving
				pointers[moving]++
				//set the next pointers to the following positions
				//next is the index of the next pointer
				next := moving + 1
				for j := pointers[moving] + 1; j < n && next < size; j++ {
					pointers[next] = j
					//if the pointer set still have not reached it's
					//max position set it as the moving pointer
					if pointers[next] != n-size+next {
						moving = next
					}
					next++
				}
				break
			}
		} else {
			//move the moving pointer to the next position
			pointers[moving]++
		}
	}
	return combinations
}

//factorial - calculates the factorial of a number n
func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

//sumMultiples - calculates the sum of the multiples of given number until the limit
//taking the sum of the mutiples of 3 to limit as example:
//the last multiple of 3 is: k = limit - (limit % 3)
//∑ = 3 + 6 + 9 + ... + k therefore:
//∑ = 3*1 + 3*2 + 3*3 + ... + 3*k
//∑ = 3 * (1 + 2 + 3 + ... + k)
//the sum is 3 times the arithmetic progession from 1 to n, where n = k/3 i.e.:
//∑ = 3 * ((1 + n) * n/2) where n is the number of terms
func sumMultiples(limit int, number int) int {
	if number == 0 {
		return 0
	}
	multiples := (limit - (limit % number)) / number
	return number * ((1 + multiples) * multiples / 2)
}

//findLcm - calculates the least common multiple for a list of values
func findLcm(values []int) int {
	result := 1
	for i := 0; i < len(values); i++ {
		result = lcm(result, values[i])
	}
	return result
}

//lcm - calculate the least common multiple
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

//findGcd - calculates the greatest common divisor for a list of values
func findGcd(values []int) int {
	result := values[0]
	for i := 1; i < len(values); i++ {
		result = gcd(result, values[i])
	}
	return result
}

//gcd - calculates the greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
