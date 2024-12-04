# loop

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

    the init statement: executed before the first iteration
    the condition expression: evaluated before every iteration
    the post statement: executed at the end of every iteration


## structure

for [ Initial Statement ] ; [ Condition ] ; [ Post Statement ] {
    [Action]
}

https://www.digitalocean.com/community/tutorials/how-to-construct-for-loops-in-go

## example 

https://go.dev/tour/flowcontrol/1


```go
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```

## حلقه

حلقه‌ها را در زبان گو به ساده‌ترین شکل ممکن و فقط با استفاده از کلید واژه for و در مدل‌های مختلف (سه‌بخشی، بی نهایت، foreach و …) می‌توان پیاده‌سازی کرد.


```go
for initialization ; condition ; counter { 
	//loop codes 
} 
```

https://book.gofarsi.ir/chapter-1/go-for/

## while

```go

package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

```

## حلقه بی‌نهایت

```go
package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
	}

	fmt.Println("this line will never execute")
}
```

```go
package main

func main() {
	for {
	}
}
```

## حلقه for-range

```go
for index, value := range slice/array {}
```

```go

package main

import "fmt"

func main() {
    letters := []string{"a", "b", "c"}

    //With index and value
    fmt.Println("Both Index and Value")
    for i, letter := range letters {
        fmt.Printf("Index: %d Value:%s\n", i, letter)
    }
    
	//Only value
    fmt.Println("\nOnly value")
    for _, letter := range letters {
        fmt.Printf("Value: %s\n", letter)
    }
}
```


## حلقه for-range برای map

```go

package main

import "fmt"

func main() {
    sample := map[string]string{
        "a": "x",
        "b": "y",
    }

    //Iterating over all keys and values
    fmt.Println("Both Key and Value")
    for k, v := range sample {
        fmt.Printf("key :%s value: %s\n", k, v)
    }

    //Iterating over only keys
    fmt.Println("\nOnly keys")
    for k := range sample {
        fmt.Printf("key :%s\n", k)
    }

    //Iterating over only values
    fmt.Println("\nOnly values")
    for _, v := range sample {
        fmt.Printf("value :%s\n", v)
    }
}
```


## string 

```go
for index, character := range string {
    //Do something with index and character
}
```

```go
package main

import "fmt"

func main() {
    sample := "a£b"

    //With index and value
    fmt.Println("Both Index and Value")
    for i, letter := range sample {
        fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
    }

    //Only value
    fmt.Println("\nOnly value")
    for _, letter := range sample {
        fmt.Printf("Value:%s\n", string(letter))
    }

    //Only index
    fmt.Println("\nOnly Index")
    for i := range sample {
        fmt.Printf("Start Index: %d\n", i)
    }
}
```


## Break

```go
package main

import "fmt"

func main() {
	sum := 0
	for {
		sum++
		if sum == 10 {
			break
		}
	}

	fmt.Println(sum)
	fmt.Println("now this line will execute")
}
```

## label

```go
package main

import "fmt"

func main() {
    letters := []string{"a", "b", "c"}

	for i := 1; i < 10; i++ {
        // define a lable with name 'second' for this loop
        second:
            for i := 2; i < 9; i++ {
                for _, l := range letters {
                    if l == "b" {
                        // break the loop with second lable
                        break second
                    }
                }
            }
	}
}
```

## continue


```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continuing loop")
			continue // break here
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
```


https://www.digitalocean.com/community/tutorials/using-break-and-continue-statements-when-working-with-loops-in-go


