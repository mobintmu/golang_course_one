package variable

import (
	"fmt"
	"math/cmplx"
	"testing"
	"unsafe"
)

func TestInt(t *testing.T) {
	var firstVar int = 10
	fmt.Println(firstVar)
	_ = firstVar
}

func TestDeclare(t *testing.T) {
	a := 120
	fmt.Println(a)
}

func TestMulti(t *testing.T) {
	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c)
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

func TestVar(t *testing.T) {
	var (
		a int    = 1
		b string = "hello"
	)

	fmt.Println(a, b)
}

func TestInitial(t *testing.T) {
	var a int
	var b float32
	var c string = ""

	fmt.Println(a, b, c)
}

func TestBool(t *testing.T) {
	var a bool = true
	fmt.Println(a)

	var b bool = false
	fmt.Println(b)
}

func TestLogic(t *testing.T) {
	var a bool = true
	var b bool = false

	fmt.Println(a && b) // true and false = false
	fmt.Println(a || b) // true or false = true
	fmt.Println(!a)     // not true = false
}

func TestIf(t *testing.T) {
	var a bool
	a = true

	if a {
		fmt.Println("a is true")
	}
}

func TestSize(t *testing.T) {
	var a int16 = 89
	var b float64 = 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d bytes", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d bytes", b, unsafe.Sizeof(b)) //type and size of b
}

func TestOverflow(t *testing.T) {
	var a int8 = 127
	fmt.Println(a)
	a += 1
	fmt.Println(a)
}

func TestFloat(t *testing.T) {
	var a float32 = 1.1
	fmt.Println(a)
	var b float64 = 2.1
	fmt.Println(b)
}

func TestCompareFloat(t *testing.T) {
	var a float64 = 10.1
	var b float64 = 10.1000001

	equalOne := EqualOne(a, b)
	equalTwo := EqualTwo(a, b)

	if !(equalOne && equalTwo) { // ! ( a && b) = !a || !b
		t.Error("EqualOne and EqualTwo should be true")
	}

}

func TestCompareInt(t *testing.T) {
	var a int = 10
	var b int = 20
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

func TestComplex(t *testing.T) {
	var a complex64 = 1 + 2i
	fmt.Println(a)
	var b complex128 = 2 + 3i
	fmt.Println(b)

	var c complex64 = complex(5, 7)

	fmt.Println(a + c)
	fmt.Println(a * c)
}

func TestVar2(t *testing.T) {
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Println(ToBe, MaxInt, z)
}

func TestString(t *testing.T) {

	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)
}

func TestRune(t *testing.T) {
	myRune := 'a'
	fmt.Println("The rune is", myRune)
}

func TestHeart(t *testing.T) {
	myRune := '♥'
	fmt.Println("The rune is", myRune)
	fmt.Printf("myRune Unicode character: %c\n", myRune)
}
