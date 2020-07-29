package diffsquares

//SquareOfSum - Calcs the square of the sum
func SquareOfSum(n int) int {
	ap := arithmeticProgression(n)
	return ap * ap
}

func arithmeticProgression(n int) int {
	return ((1 + n) * n) / 2
}

//SumOfSquares - Calcs the sum of the squares
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

//Difference - difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	a := SquareOfSum(n)
	b := SumOfSquares(n)
	return a - b
}
