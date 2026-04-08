package godash

import "math"

// IsWhole returns true if v is whole number
func IsWhole(v float64) bool {
	return math.Abs(math.Remainder(v, 1)) == 0
}
