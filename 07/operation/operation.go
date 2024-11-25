package operation

import (
	"errors"
	"fmt"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func swap(first int32, second int32) (int32, int32) {
	return second, first
}

func example(a int) {
	if a == 10 {
		return
	}
	fmt.Println(a)
}

func recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursive(n-1)
}

func multiply(a, b int) int {
	a = a * 2 // modifying a inside the function
	return a * b
}

func multiplyPointer(a, b *int) int {
	*a = *a * 2 // modifying a's value at its memory address
	return *a * *b
}

func ContainA(input string) (bool, error) {
	if input == "" {
		return false, fmt.Errorf("Empty input")
	}
	res := strings.Contains(input, "a")
	return res, nil
}

func DividedZero(max int) error {
	for i := 0; i < max; i++ {
		if i == 0 {
			return errors.New("Divided by zero")
		}
		result := 20 / i
		fmt.Println(result)
	}
	return nil
}

func calculateDivide(a, b int) int {
	if b == 0 {
		panic("division by zero occurred")
	}
	return a / b
}
