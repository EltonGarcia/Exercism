package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
