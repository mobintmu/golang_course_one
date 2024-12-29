package solve

import (
	"fmt"
	"math"
)

func CheckEquation(x, y int64) (float64, bool) {
	value := math.Pow(float64(x), 2) + math.Pow(float64(y), 2)
	return value, value == 40501
}

type Solution struct {
	X, Y int64
}

func Find() []Solution {
	var i, j float64
	var solutions []Solution
	// O(n^2)
	for i = -1000; i < 1000; i += 0.0001 {
		for j = -1000; j < 1000; j += 0.0001 {
			fmt.Println(i, j)
			value, res := CheckEquation(int64(i), int64(j))
			_ = value
			if res {
				solutions = append(solutions, Solution{int64(i), int64(j)})
			}
		}
	}
	// O(n^2)
	return FindDuplicatSolutions(solutions)
}

func FindDuplicatSolutions(input []Solution) []Solution {
	result := []Solution{}
	for _, element := range input {
		key := false
		for _, res := range result {
			if element.X == res.X && element.Y == res.Y || element.X == res.Y && element.Y == res.X {
				key = true
				break
			}
		}
		if !key {
			result = append(result, element)
		}
	}
	return result
}
