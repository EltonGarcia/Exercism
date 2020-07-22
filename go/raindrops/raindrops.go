package raindrops

import (
	"strconv"
)

//Convert to raindrops
func Convert(n int) string {
	switch 0 {
	case n % 105:
		return "PlingPlangPlong"
	case n % 35:
		return "PlangPlong"
	case n % 21:
		return "PlingPlong"
	case n % 15:
		return "PlingPlang"
	case n % 3:
		return "Pling"
	case n % 5:
		return "Plang"
	case n % 7:
		return "Plong"
	default:
		return strconv.Itoa(n)
	}
}
