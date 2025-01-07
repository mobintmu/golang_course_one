package pointer

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	res := calculate(add, 5, 3)
	fmt.Println(res)

	res = calculate(subtract, 5, 3)
	fmt.Println(res)

}
