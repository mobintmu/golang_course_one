package loop

import "fmt"

func Sum(maxInput int) int {
	sum := 0
	for counter := 1; counter < maxInput; counter++ {
		fmt.Println("counter: ", counter)
		sum += counter
		fmt.Println("sum: ", sum)

	}
	fmt.Println("Sum:", sum)
	return sum
}

func SumWithpoutInitial(min, max int) int {
	sum := 0
	for ; min < max; min++ {
		sum += min
	}
	return sum
}

func Infinite(min, max int) int {
	sum := 0
	for {
		// sum += min
		fmt.Println(sum)
	}
	return sum
}

func forrange(input []int) {
	for _, value := range input {
		fmt.Println(value)
	}
}

func forMap(input map[int]int) {
	for key, value := range input {
		fmt.Println(key, value)
	}
}

func forString(input string) {
	for _, value := range input {
		fmt.Println(string(value))
	}
}

func SumCondition(min, max int) int {
	sum := 0
	for ; min < max; min++ {
		if min == 3 {
			break
		}
		sum += min
	}
	return sum
}

func SumContinue(min, max int) int {
	sum := 0
	for ; min < max; min++ {
		if min%2 == 0 {
			continue
		}
		sum += min
	}
	return sum
}
