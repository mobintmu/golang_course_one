# Go Variable

## Go Variable Types

1. int- stores integers (whole numbers), such as 123 or -123
2. float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
3. string - stores text, such as "Hello World". String values are surrounded by double quotes
4. bool- stores values with two states: true or false

https://www.w3schools.com/go/go_variables.php


## Debug mode

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Go Debug",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/.",
            "cwd": "${workspaceFolder}/"
        }
    ]
}
```


## Declaring (Creating) Variables

```go
var variablename type = value
var myVariable int = 10

```

## With the := sign:

```go
 variablename := value 
```

## Scope

```go
package main
import ("fmt")

var a int
var b int = 2
var c = 3

func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

## Since := is used outside of a function, running the program results in an error. 

```go
package main
import ("fmt")

a := 1

func main() {
  fmt.Println(a)
}
```

## Go Multiple Variable Declaration

```go
package main
import ("fmt")

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
```


## Go Variable Declaration in a Block

```go
package main
import ("fmt")

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

## Go naming conventions for const

The standard library uses camel-case, so I advise you do that as well. The first letter is uppercase or lowercase depending on whether you want to export the constant.

A few examples:

    md5.BlockSize
    os.O_RDONLY is an exception because it was borrowed directly from POSIX.
    os.PathSeparator

https://stackoverflow.com/questions/22688906/go-naming-conventions-for-const


## Golang Variables, Zero Values, and Type inference


![variable](https://d33wubrfki0l68.cloudfront.net/5e55d8febfacbe21edb64cfe7bcc5248532358df/02c3d/static/7a5355bf9eaecbbf401f7637654f0508/a76f4/golang-variables-data-types-illustration.png)


```go
var firstName string
var lastName string  
var age int
```


## bool

bool type represents a boolean. It can either be a true or false value.

```go
package main
import "fmt"
func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}
```

## Signed integers

- int8 -128 to 127
- int16 -32768 to 32767
- int32 -2147483648 to 2147483647
- int64 -9223372036854775808 to 9223372036854775807
- int -9223372036854775808 to 9223372036854775807
- uint8 0 to 255
- uint16 0 to 65535
- uint32 0 to 4294967295
- uint64 0 to 18446744073709551615
- uint 0 to 18446744073709551615

```go
package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var a = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d bytes", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d bytes", b, unsafe.Sizeof(b)) //type and size of b
}
```

https://golangbot.com/types/


## Floating point types

- float32 32 bit floating point numbers
- float64 64 bit floating point numbers

```go
package main
import (
	"fmt"
)
func main() {
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Printf("sum of %f and %f is %f, diff is %f\n", a, b, sum, diff)
	no1, no2 := 56, 89
	fmt.Printf("type of no1 %T no2 %T\n", no1, no2)
	fmt.Printf("sum of %d and %d is %d, diff is %d", no1, no2, no1+no2, no1-no2)
}
```


### Go float comparison

```go
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}

func main() {
    a := 0.1
    b := 0.2
    fmt.Println(almostEqual(a + b, 0.3))
}
```

https://stackoverflow.com/questions/47969385/go-float-comparison

## Complex types

- complex64 complex numbers which have float32 real and imaginary parts
- complex128 complex numbers which have float64 real and imaginary parts

```go
package main
import (
	"fmt"
)
func main() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}
```

https://byjus.com/maths/complex-numbers/


```go
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

https://go.dev/tour/basics/11


## String types

- string represents a sequence of bytes
- string is a collection of characters enclosed in double quotes
- string is immutable
- A string represents a sequence of characters. Strings are immutable; once they are created they can’t be modified.

```go
package main
import (
    "fmt"
)
func main() {
    first := "Naveen"
    last := "Ramanathan"
    name := first + " " + last
    fmt.Println("My name is", name)
}
```

https://www.codecademy.com/resources/docs/go/data-types

## Rune type

Used for characters. Internally used as 32-bit integers.

The rune type in Go is an alias for int32 . Given this underlying int32 type, the rune type holds a signed 32-bit integer value. However, unlike an int32 type, the integer value stored in a rune type represents a single Unicode character.

Unicode and Unicode Code Points

Unicode is a superset of ASCII that represents characters by assigning a unique number to every character. This unique number is called a Unicode code point. Unicode aims to represent all the world's characters including various alphabets, numbers, symbols, and even emoji as Unicode code points.

In Go, the rune type represents a single Unicode code point.

https://exercism.org/tracks/go/concepts/runes

```go
package main
import (
    "fmt"
)
func main() {
    myRune := 'a'
    fmt.Println("The rune is", myRune)
}
```

```go
package main
import (
    "fmt"
)
func main() {
    myRune := '♥'
    fmt.Println("The rune is", myRune)
    fmt.Printf("myRune Unicode character: %c\n", myRune)

    myRune := rune(0xbf)
    myRune = 191
    fmt.Printf("myRune Unicode character: %c\n", myRune)
}
```

```go
myRuneSlice := []rune{'e', 'x', 'e', 'r', 'c', 'i', 's', 'm'}
myString := string(myRuneSlice)
fmt.Println(myString)
```


## memory

In Go, variables are stored in memory based on their type and scope, managed by the Go runtime. 
Here is a broad overview of how it works:

Memory Allocation

- Stack Memory: Used for local variables within functions. 
  It's fast and automatically managed. Once the function call is over, the memory is freed.

- Heap Memory: Used for dynamic memory allocation, such as variables created with new or make.
  The garbage collector manages it, cleaning up unused objects.

Stack Allocation:
```go
func main() {
    a := 10  // Stored on the stack
    b := 20  // Stored on the stack
    fmt.Println(a, b)
}
```


Heap Allocation:
```go
func main() {
    c := new(int)  // Allocates memory on the heap
    *c = 30
    fmt.Println(*c)
}
```

Global Variables: Stored in the data segment and exist for the program’s lifetime.

```go
var globalVar int = 100 // Global scope
```

Garbage Collection

Go has an automatic garbage collector that frees up memory that is no longer in use, 
helping manage heap-allocated variables efficiently.