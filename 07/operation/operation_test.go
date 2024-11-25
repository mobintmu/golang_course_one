package operation

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestSwap(t *testing.T) {
	first, second := swap(1, 2)
	if first != 2 || second != 1 {
		t.Errorf("Expected 2, 1 but got %d, %d", first, second)
	}
}

func TestExample(t *testing.T) {
	example(10)
	// This is a dummy test, it is just to show that we can test a function that does not return anything
	// and does not have any side effects
}

func TestRecursive(t *testing.T) {
	result := recursive(10)
	fmt.Println(result)
}

func TestMultiply(t *testing.T) {
	a := 2
	b := 3
	result := multiply(a, b)
	fmt.Println(result)
	fmt.Println("a :", a, "b : ,", b)

	result = multiplyPointer(&a, &b)
	fmt.Println(result)
	fmt.Println("a :", a, "b : ,", b)
}

func TestContain(t *testing.T) {
	result, err := ContainA("")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !result {
		t.Errorf("Expected true, but got %v", result)
	}
}

func TestDividedZero(t *testing.T) {
	err := DividedZero(10)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCalculateDivide(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic:", r)
		}
	}()
	result := calculateDivide(10, 0)
	if result != 0 {
		t.Errorf("Expected 0, but got %d", result)
	}

	fmt.Println("Finish")
}
