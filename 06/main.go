package main

import "fmt"

var MyGlobal int = 200

func main() {
	var a int = -100
	fmt.Println(a)
	fmt.Println(MyGlobal)
	{
		var a float32 = 1.1
		fmt.Println(a)
	}

	fmt.Println(a)

	myFunc()
}

func myFunc() {
	fmt.Println(MyGlobal)
	var a int = 2000
	fmt.Println(a)
}
