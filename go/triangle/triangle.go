package triangle

import (
	"math"
)

// Kind of triangle
type Kind int

const (
	//NaT - not a triangle
	NaT = iota
	//Equ equilateral triangle
	Equ
	//Iso isosceles triangle
	Iso
	//Sca scalene triangle
	Sca
)

func isTriangleValid(a, b, c float64) bool {
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	return math.Abs(a-b) <= c && c <= a+b
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch true {
	case math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) || !isTriangleValid(a, b, c):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}
