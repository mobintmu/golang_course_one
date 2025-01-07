package pointer

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestInitialPointer(t *testing.T) {

	var one int = 32
	fmt.Println(one)
	fmt.Println("_________________________________")

	var two *int = &one
	fmt.Println(two)
	fmt.Println(*two)
	fmt.Println("_________________________________")

	*two = 22
	fmt.Println(one)
}

func TestInitial(t *testing.T) {
	var first Point = Point{1, 2}
	fmt.Println(first)

	var second *Point = &first
	fmt.Println(*second)

	second.x = 3
	fmt.Println(first)
}

func TestArray(t *testing.T) {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	ptr := &arr[0]

	for i := 0; i < len(arr); i++ {
		ptr = &arr[i]
		fmt.Printf("%d ", *ptr)
		fmt.Println(ptr)
	}
}

func TestNil(t *testing.T) {
	var ptr *Point
	fmt.Println(ptr)
	ptr = &Point{2, 3}
	if ptr == nil {
		fmt.Println("ptr is nil")
		return
	}
	fmt.Println(*ptr)
}

func TestPointerForPointer(t *testing.T) {
	var valid int = 10

	var ptr *int = &valid
	fmt.Println(ptr) //0xc000012380
	fmt.Println(*ptr)

	var ptrPtr **int = &ptr
	fmt.Println(ptrPtr)   // 0xc00007c070
	fmt.Println(*ptrPtr)  //0xc000012380
	fmt.Println(**ptrPtr) // 10

}

func TestNew(t *testing.T) {
	var ptr *int = new(int)
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	Change(a, b)
	fmt.Println(a, b)
}

func TestAddPointer(t *testing.T) {
	a := 1
	b := 2
	ChangePointer(&a, &b)
	fmt.Println(a, b)
}

func TestNewPoint(t *testing.T) {
	p, err := NewPoint(-1, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}

func TestPTR(t *testing.T) {
	arr := [3]int{1, 2, 3}
	ptr := &arr[0]
	fmt.Println(*ptr) // output: 1

	ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
	fmt.Println(*ptr) // output: 2

	ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
	fmt.Println(*ptr) // output: 3

	ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
	fmt.Println(*ptr) // output: ?

}
