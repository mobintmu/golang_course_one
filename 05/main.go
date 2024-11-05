package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	// fmt.Println(Hello())

	c, python, java := true, false, "no!"

	fmt.Println(c, python, java)

}
