package hamming

import (
	"errors"
)

//Distance - Calcs Hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		err := errors.New("Different sizes")
		return -1, err
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
