package solve

import (
	"fmt"
	"testing"
)

func TestEquation(t *testing.T) {
	value, result := CheckEquation(1, 2)
	fmt.Println(value)
	if result {
		fmt.Println("find")
	}
	t.Error("Expected 40501, but got", value)
}

func TestFind(t *testing.T) {
	fmt.Println("Start")
	solutions := Find()
	if len(solutions) == 0 {
		t.Error("No solution found")
	}
	for _, solution := range solutions {
		fmt.Println("____________________")
		fmt.Println(solution.X, solution.Y)
	}
}
