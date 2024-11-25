# Function in Golang


## Functions
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.


```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

## Single responsibility principle


The single-responsibility principle (SRP) is a computer programming principle that states that "A module should be responsible to one, and only one, actor."[1] The term actor refers to a group (consisting of one or more stakeholders or users) that requires a change in the module.

Robert C. Martin, the originator of the term, expresses the principle as, "A class should have only one reason to change".[2] Because of confusion around the word "reason", he later clarified his meaning in a blog post titled "The Single Responsibility Principle", in which he mentioned Separation of Concerns and stated that "Another wording for the Single Responsibility Principle is: Gather together the things that change for the same reasons. Separate those things that change for different reasons.".[3] In some of his talks, he also argues that the principle is, in particular, about roles or actors. For example, while they might be the same person, the role of an accountant is different from a database administrator. Hence, each module should be responsible for each role.[4]

https://en.wikipedia.org/wiki/Single-responsibility_principle


## Swap two variables


```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

https://go.dev/tour/basics/6


### Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

```go
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

https://go.dev/tour/basics/7


## Go Recursion Functions

Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. We use the x variable as the data, which increments with 1 (x + 1) every time we recurse. The recursion ends when the x variable equals to 11 (x == 11). 


```go
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
```


Recursion is a common mathematical and programming concept. This has the benefit of meaning that you can loop through data to reach a result.


```go
package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}
```
https://www.w3schools.com/go/go_functions.php



## Call by Value



In call by value, values of the arguments are copied to the function parameters, so changes in the function do not affect the original variables.

```go
package main
import "fmt"

func multiply(a, b int) int {
    a = a * 2 // modifying a inside the function
    return a * b
}

func main() {
    x := 5
    y := 10
    fmt.Printf("Before: x = %d, y = %d\n", x, y)
    result := multiply(x, y)
    fmt.Printf("multiplication: %d\n", result)
    fmt.Printf("After: x = %d, y = %d\n", x, y)
}
```


https://www.geeksforgeeks.org/functions-in-go-language/


## Call by Reference

package main
import "fmt"

func multiply(a, b *int) int {
    *a = *a * 2 // modifying a's value at its memory address
    return *a * *b
}

func main() {
    x := 5
    y := 10
    fmt.Printf("Before: x = %d, y = %d\n", x, y)
    result := multiply(&x, &y)
    fmt.Printf("multiplication: %d\n", result)
    fmt.Printf("After: x = %d, y = %d\n", x, y)
}


## Return and handle an error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings module, then handle it in the caller.


```go
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
```

https://go.dev/doc/tutorial/handle-errors



## Go Errors

When we run this code, we will get an error called integer divide by zero.



```go
package main
import "fmt"

func main() {

  for i := 0; i < 5; i++ {
    result := 20 / i
    fmt.Println(result)
  }
}
```

### Go Error using New() Function


```go
package main

// import the errors package
import (
  "errors"
  "fmt"
)

func main() {

  message := "Hello"

  // create error using New() function
  myError := errors.New("WRONG MESSAGE")

if message != "Programiz" {
  fmt.Println(myError)
}
  
}
```

### Example: Error using New()

```go
package main

// import the errors package
import (
  "errors"
  "fmt"
)

// function that checks if name is Programiz
func checkName(name string) error {

  // create a new error
  newError := errors.New("Invalid Name")

  // return the error if name is not Programiz
  if name != "Programiz" {
    return newError
  }

  // return nil if there is no error
  return nil
}

func main() {

  name := "Hello"

  // call the function
  err := checkName(name)

  // check if the err is nil or not
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Valid Name")
  }

}
```
https://www.programiz.com/golang/errors



## Go by Example: Errors


In Go it’s idiomatic to communicate errors via an explicit, separate return value. This contrasts with the exceptions used in languages like Java and Ruby and the overloaded single result / error value sometimes used in C. Go’s approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks.

See the documentation of the errors package and this blog post for additional details.


```go

package main

import (
    "errors"
    "fmt"
)

func f(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}

func main() {
    for _, i := range []int{7, 42} {

        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }

    for i := range 5 {
        if err := makeTea(i); err != nil {

            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }

        fmt.Println("Tea is ready!")
    }
}

```


https://gobyexample.com/errors



## Error handling in Golang


### Displaying errors manually

```go
package main

import (
 “fmt”
 “errors”

 
)

func division(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New(“b cannot be 0”)
    }
 return a/b, nil
}

    

func main() {
 result, err := division(10, 3)
    if err != nil {
        fmt.Println(“Error:”, err)
    } else {
        fmt.Println(“Result:”, result)
    }

}
```

### Creating a custom error using fmt.Errorf


```go
func makeDivision(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf(“cannot divide %d by %d”, a, b)
    }
 return a/b, nil
}
```

### Using Panic and Recover

```go
package main

import (
 “fmt”
 )

func calculateDivide(a, b int) int {
 if b == 0 {
  panic(“division by zero occurred”)
 }
 return a/b
}

func main() {
 
 fmt.Println(“Start program”)
    result := calculateDivide(10, 0) 
    fmt.Println(“Result:”, result)
    fmt.Println(“End of program”) 

}
```

### Using defer to recover from panic

```go
package main

import (
    “fmt”
)

func calculateDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(“Recover from panic:”, r)
        }
    }()
    fmt.Println(“Division result:”, a/b)
}

func main() {
    fmt.Println(“Start program”)
    calculateDivide(10, 0) 
    fmt.Println(“End of program”) 
}
```

https://medium.com/@nandoseptian/error-handling-in-golang-418a73e19229



