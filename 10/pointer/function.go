package pointer

type ArithFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func calculate(op ArithFunc, a, b int) int {
	return op(a, b)
}
