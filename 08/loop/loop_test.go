package loop

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(6)
	if sum != 15 {
		t.Error("Expected 15, but got", sum)
	}
}

func TestSumWithoutInitial(t *testing.T) {
	sum := SumWithpoutInitial(2, 6)
	if sum != 14 {
		t.Error("Expected 14, but got", sum)
	}
}

func TestInfinite(t *testing.T) {
	sum := Infinite(0, 10)
	t.Log("Infinite loop sum:", sum)
}

func TestForrange(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	forrange(input)
}

func TestForMap(t *testing.T) {
	input := map[int]int{1: 10, 2: 20, 3: 30}
	forMap(input)
}

func TestForString(t *testing.T) {
	input := "Hello, World!"
	forString(input)
}

func TestSumCondit(t *testing.T) {
	sum := SumCondition(0, 10)
	if sum != 3 {
		t.Error("Expected 3, but got", sum)
	}
}

func TestSumContinue(t *testing.T) {
	sum := SumContinue(0, 10)
	// 1 + 3 + 5 + 7 + 9 = 16
	if sum != 16 {
		t.Error("Expected 25, but got", sum)
	}
}
