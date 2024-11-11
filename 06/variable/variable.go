package variable

import (
	"math"
)

func EqualOne(a float64, b float64) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func EqualTwo(a float64, b float64) bool {
	if math.Abs(float64(a-b)) < 1e-5 {
		return true
	} else {
		return false
	}
}
