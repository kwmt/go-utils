package utils

// Convert to percent two ints.
func Percent(x int, y int) float64 {
	return Div(x, y) * 100
}

// Divide x by y
func Div(x int, y int) float64 {
	return float64(x) / float64(y)
}
