package utils

import "math"

func isSquare(N float64) bool {
	n := math.Sqrt(N)
	return n == float64(int(n))
}
