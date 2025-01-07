# pointer

 Go has pointers. A pointer holds the memory address of a value. 

 https://go.dev/tour/moretypes/1


Address-of Operator (&): This operator returns the memory address of a variable.
Dereference Operator (*): This operator accesses the value stored at the memory address pointed to by a pointer.



 ```go
 package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

```go
x := 10
ptr := &x
fmt.Println(*ptr) // Prints 10
*ptr = 20
fmt.Println(x)    // Prints 20
```

```go
arr := [5]int{1, 2, 3, 4, 5}
ptr := &arr[0]
for i := 0; i < len(arr); i++ {
    fmt.Printf("%d ", *ptr)
    ptr++
}
```

## nil

```go
var ptr *int
if ptr == nil {
    fmt.Println("Pointer is nil")
}
```

## memory

https://www.geeksforgeeks.org/pointers-in-golang/

![pic](https://media.geeksforgeeks.org/wp-content/uploads/20190705160332/Pointers-in-Golang.jpg)




## A Comprehensive Guide to Pointers in Go

https://medium.com/@jamal.kaksouri/a-comprehensive-guide-to-pointers-in-go-4acc58eb1f4d


Memory Management with Pointers:
```go
var ptr *int = new(int)

fmt.Println(ptr) // output: 0xc0000160c0
fmt.Println(*ptr) // output: 0

*ptr = 10
fmt.Println(*ptr) // output: 10

ptr = nil
```

Pass by Value vs. Pass by Reference:
```go
func addOne(x *int) {
  *x++
}

func main() {
  x := 10
  fmt.Println(x) // output: 10
  addOne(&x)
  fmt.Println(x) // output: 11
}
```

Pointer Arithmetic:
```go
func main() {
  arr := [3]int{1, 2, 3}
  ptr := &arr[0]
  fmt.Println(*ptr) // output: 1
  ptr++
  fmt.Println(*ptr) // output: 2
}
```

Note: Please exercise caution when using the unsafe package, as it bypasses some of Go's safety mechanisms. It's generally recommended to avoid using unsafe unless absolutely necessary and to explore alternative approaches whenever possible.

Here’s an example of how you can increment a pointer using the unsafe package:
```go
import (
 "fmt"
 "unsafe"
)

func main() {
 arr := [3]int{1, 2, 3}
 ptr := &arr[0]
 fmt.Println(*ptr) // output: 1

 ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
 fmt.Println(*ptr) // output: 2
}
```

Arrays and Pointers:
```go
func printArray(arr *[3]int) {
  fmt.Println(*arr)
}

func main() {
  arr := [3]int{1, 2, 3}
  ptr := &arr
  fmt.Println(*ptr) // output: [1 2 3]
  printArray(ptr) // output: [1 2 3]
}
```

Functions and Pointers:

```go
func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func double(x *int) *int {
  result := *x * 2
  return &result
}

func main() {
  x := 1
  y := 2
  fmt.Println(x, y) // output: 1 2
  swap(&x, &y)
  fmt.Println(x, y) // output: 2 1
  ptr := double(&x)
  fmt.Println(*ptr) // output: 4
}
```

Function Types in Go:
```go
package main

import "fmt"

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

func main() {
    fmt.Println(calculate(add, 5, 3))  // Prints: 8
    fmt.Println(calculate(subtract, 5, 3))  // Prints: 2
}
```




