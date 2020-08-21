package grains

import (
	"errors"
)

//Square - Calculates the number of grains on a square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("N must be greater than one and less than 64")
	}
	return 1 << (n - 1), nil
}

//Total - Calculates the total of grains in the chessboard
//Using Telescoping series we find that the sum of 2^x from 0 to n is (2^(n+1)) - 1
func Total() uint64 {
	return 1<<64 - 1
}
